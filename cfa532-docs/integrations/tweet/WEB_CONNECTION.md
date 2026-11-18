# Leither Connection Management

**Canonical home:** `cfa532/Leither/cfa532-docs/`
**Consolidated:** 2026-10-16 from [TweetWeb/docs/LEITHER_CONNECTION.md](https://github.com/cfa532/TweetWeb/blob/6246c715de6de873630d2902a96862972d2cd26a/docs/LEITHER_CONNECTION.md).

Historical TweetWeb integration note. Code samples and performance claims are preserved from the source document; this move does not revalidate them against current implementation.

---

This document describes the centralized leither connection management system implemented in TweetWeb.

## Overview

The leither connection management has been centralized in `app.js` to provide better performance, consistency, and maintainability across all services that require IPFS operations.

## Architecture

### Centralized Management

All leither connections are managed from a single location (`app.js`) and shared across all service modules:

- **Video Processing** (`videoRoutes.js`)
- **ZIP Processing** (`zipRoutes.js`)
- **Future Services** (any new modules requiring IPFS access)

### Key Benefits

1. **Performance**: Leither port detected once at startup, cached in memory
2. **Efficiency**: Shared connection pool across all services
3. **Consistency**: All modules use identical connection patterns
4. **Maintainability**: Single point of control for connection logic
5. **Reliability**: Centralized error handling and connection management

## Implementation

### Startup Initialization

```javascript
// In app.js - runs once when server starts
async function initializeLeither() {
  if (leitherInitialized) {
    return;
  }

  try {
    console.log('[LEITHER-INIT] Initializing Leither connection...');
    leitherPort = await getLeitherPort(); // Uses process name detection for accuracy
    console.log(`[LEITHER-INIT] Detected Leither service on port: ${leitherPort}`);

    // Test connection by creating a client
    const testClient = await createLeitherClient(leitherPort);
    if (testClient) {
      console.log('[LEITHER-INIT] Leither connection test successful');
      leitherInitialized = true;
    }
  } catch (error) {
    console.error('[LEITHER-INIT] Failed to initialize Leither connection:', error.message);
    console.log('[LEITHER-INIT] Leither service may not be available, but app will continue');
    leitherInitialized = false;
  }
}

// Called during server startup
app.listen(port, '::', async () => {
  console.log(`Server listening on port ${port}`);
  // Initialize Leither connection on startup
  await initializeLeither();
});
```

### Connection Pool Management

```javascript
// Global connection pool
let leitherConnections = new Map();
let leitherPort = null;
let leitherInitialized = false;
const maxLeitherConnections = 2;

// Get or create a connection from the pool
async function getLeitherConnection() {
  if (!leitherInitialized || !leitherPort) {
    throw new Error('Leither service not initialized');
  }

  const connectionKey = `port-${leitherPort}`;

  // Check for available connection
  if (leitherConnections.has(connectionKey)) {
    const connection = leitherConnections.get(connectionKey);
    if (connection.isAvailable) {
      connection.isAvailable = false;
      return connection.client;
    }
  }

  // Create new connection if under limit
  if (leitherConnections.size < maxLeitherConnections) {
    const client = await createLeitherClient(leitherPort);
    leitherConnections.set(connectionKey, {
      client,
      isAvailable: false,
      port: leitherPort
    });
    return client;
  }

  // Wait for available connection
  return new Promise((resolve) => {
    const checkConnection = () => {
      for (const [key, conn] of leitherConnections) {
        if (conn.isAvailable) {
          conn.isAvailable = false;
          resolve(conn.client);
          return;
        }
      }
      setTimeout(checkConnection, 100);
    };
    checkConnection();
  });
}

// Release connection back to pool
function releaseLeitherConnection(client) {
  for (const [key, conn] of leitherConnections) {
    if (conn.client === client) {
      conn.isAvailable = true;
      break;
    }
  }
}
```

### Global Function Exposure

```javascript
// Make leither functions available globally
global.getLeitherConnection = getLeitherConnection;
global.releaseLeitherConnection = releaseLeitherConnection;
global.getCurrentLeitherPort = getCurrentLeitherPort;
global.isLeitherInitialized = isLeitherInitialized;
```

## Usage in Service Modules

### Video Processing (videoRoutes.js)

```javascript
// Before (old implementation)
const leitherPort = await getLeitherPort(); // Port detection on every request
const leitherClient = await getLeitherConnection(); // Local connection management

// After (centralized implementation)
const leitherPort = global.getCurrentLeitherPort(); // Use cached port
const leitherClient = await global.getLeitherConnection(); // Use shared pool

// Always release connection when done
global.releaseLeitherConnection(leitherClient);
```

### ZIP Processing (zipRoutes.js)

```javascript
// Same pattern as video processing
const leitherPort = global.getCurrentLeitherPort();
const leitherClient = await global.getLeitherConnection();

// Process with leither...
const cid = await leitherClient.IpfsAdd(api.sid, hlsContentPath);

// Release connection
global.releaseLeitherConnection(leitherClient);
```

## Connection Lifecycle

### 1. Server Startup
```
┌─────────────────┐
│   Server Start  │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ initializeLeither() │
│ ├── Detect port │
│ ├── Test conn   │
│ └── Cache info  │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ Ready to serve  │
└─────────────────┘
```

### 2. Request Processing
```
┌─────────────────┐
│ Service Request │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ getLeitherConn() │
│ ├── Check pool  │
│ ├── Get/reuse   │
│ └── Mark busy   │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ Process Request │
└─────────┬───────┘
          │
          ▼
┌─────────────────┐
│ releaseLeither() │
│ └── Mark avail  │
└─────────────────┘
```

## Performance Improvements

### Before Centralization

- **Port Detection**: Every request detected leither port (1-2 seconds)
- **Connection Overhead**: Each service created its own connections
- **Resource Waste**: Duplicate connection pools
- **Inconsistent Behavior**: Different connection patterns across services

### After Centralization

- **Port Detection**: Once at startup (cached in memory)
- **Connection Reuse**: Shared pool across all services
- **Resource Efficiency**: Single connection pool (max 2 connections)
- **Consistent Behavior**: Identical connection patterns everywhere

### Performance Metrics

| Metric | Before | After | Improvement |
|--------|--------|-------|-------------|
| Port Detection | Per request | Once at startup | ~95% reduction |
| Connection Setup | Per service | Shared pool | ~80% reduction |
| Memory Usage | Duplicate pools | Single pool | ~50% reduction |
| Request Latency | +1-2s port detection | Immediate | ~1-2s faster |

## Error Handling

### Connection Failures

```javascript
try {
  const leitherClient = await global.getLeitherConnection();
  // Process request...
} catch (error) {
  if (error.message === 'Leither service not initialized') {
    // Handle initialization failure
    console.error('Leither service not available');
    return res.status(503).json({ error: 'Service temporarily unavailable' });
  }
  // Handle other connection errors
  throw error;
} finally {
  if (leitherClient) {
    global.releaseLeitherConnection(leitherClient);
  }
}
```

### Startup Failures

```javascript
// In initializeLeither()
try {
  leitherPort = await getLeitherPort();
  const testClient = await createLeitherClient(leitherPort);
  leitherInitialized = true;
} catch (error) {
  console.error('[LEITHER-INIT] Failed to initialize:', error.message);
  leitherInitialized = false;
  // Server continues to run, but leither-dependent features will fail gracefully
}
```

## Monitoring and Debugging

### Logging

The system provides comprehensive logging:

```javascript
// Startup logs
[LEITHER-INIT] Initializing Leither connection...
[LEITHER-INIT] Detected Leither service on port: 8081
[LEITHER-INIT] Leither connection test successful

// Connection logs
[LEITHER] Creating client connection to port 8081
[LEITHER] WebSocket connection established to port 8081

// Service logs
[INFO] Using Leither service on port: 8081
[SUCCESS] IPFS CID received: QmXXX...
```

### Health Checks

```javascript
// Check if leither is initialized
if (global.isLeitherInitialized()) {
  console.log('Leither service is available');
} else {
  console.log('Leither service is not available');
}

// Get current port
const port = global.getCurrentLeitherPort();
console.log(`Leither running on port: ${port}`);
```

## Migration Guide

### For Existing Services

1. **Remove local leither code**:
   ```javascript
   // Remove these from service files
   const leitherConnections = new Map();
   async function getLeitherConnection() { ... }
   function releaseLeitherConnection() { ... }
   ```

2. **Update function calls**:
   ```javascript
   // Change from
   const leitherPort = await getLeitherPort();
   const leitherClient = await getLeitherConnection();

   // To
   const leitherPort = global.getCurrentLeitherPort();
   const leitherClient = await global.getLeitherConnection();
   ```

3. **Add connection release**:
   ```javascript
   // Always release connections
   try {
     // Process with leither...
   } finally {
     if (leitherClient) {
       global.releaseLeitherConnection(leitherClient);
     }
   }
   ```

### For New Services

1. **Use global functions**:
   ```javascript
   const leitherClient = await global.getLeitherConnection();
   // Process request...
   global.releaseLeitherConnection(leitherClient);
   ```

2. **Handle initialization errors**:
   ```javascript
   if (!global.isLeitherInitialized()) {
     throw new Error('Leither service not available');
   }
   ```

## Best Practices

1. **Always release connections**: Use try/finally blocks
2. **Check initialization**: Verify leither is available before use
3. **Handle errors gracefully**: Provide meaningful error messages
4. **Monitor performance**: Watch for connection pool exhaustion
5. **Test thoroughly**: Ensure all services work with centralized management

## Troubleshooting

### Common Issues

1. **"Leither service not initialized"**
   - Check if leither service is running
   - Verify port detection in logs
   - Check network connectivity

2. **Connection pool exhaustion**
   - Monitor active connections
   - Ensure connections are being released
   - Consider increasing pool size if needed

3. **Port detection failures**
   - Check leitherDetector.js logs
   - Verify Leither process is running: `lsof -i -P -n | grep LISTEN | grep Leither`
   - Verify detection commands work on your OS:
     - **macOS**: `lsof -i -P -n | grep LISTEN` or `netstat -an | grep LISTEN`
     - **Linux**: `netstat -tlnp | grep LISTEN` or `lsof -i -P -n | grep LISTEN`
   - Check firewall settings
   - Ensure leither service is running on ports 8000-9000
   - **Note**: Detection now uses process name "Leither" (case-sensitive) for accuracy

### Debug Commands

```bash
# Check leither service status
curl -s http://localhost:8081/webapi/

# Monitor server logs
tail -f /var/log/tweetweb/server.log | grep LEITHER

# Check for Leither process specifically (most accurate)
lsof -i -P -n | grep LISTEN | grep Leither

# Check listening ports (cross-platform)
lsof -i -P -n | grep LISTEN | grep -E ":(80[0-9][0-9]|90[0-9][0-9])"

# macOS specific
netstat -an | grep LISTEN | grep -E "\.(80[0-9][0-9]|90[0-9][0-9])"

# Linux specific
netstat -tlnp | grep LISTEN | grep -E ":(80[0-9][0-9]|90[0-9][0-9])"

# Check connection pool status
# (Add monitoring endpoint to app.js if needed)
```

## Future Enhancements

1. **Connection Health Monitoring**: Periodic health checks
2. **Dynamic Pool Sizing**: Adjust pool size based on load
3. **Connection Retry Logic**: Automatic reconnection on failures
4. **Metrics Collection**: Connection usage statistics
5. **Load Balancing**: Multiple leither service support

## Related Documentation

- [Architecture Overview](https://github.com/cfa532/TweetWeb/blob/6246c715de6de873630d2902a96862972d2cd26a/docs/ARCHITECTURE.md)
- [Video Conversion Guide](https://github.com/cfa532/TweetWeb/blob/6246c715de6de873630d2902a96862972d2cd26a/docs/VIDEO_CONVERSION.md)
- [ZIP Processing Guide](https://github.com/cfa532/TweetWeb/blob/6246c715de6de873630d2902a96862972d2cd26a/docs/ZIP_PROCESSING.md)
- [API Documentation](https://github.com/cfa532/TweetWeb/blob/6246c715de6de873630d2902a96862972d2cd26a/docs/API.md)
