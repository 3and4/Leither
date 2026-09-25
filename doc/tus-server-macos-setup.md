# Setting Up tus-server as Agents on macOS

This guide explains how to configure tus-server to run as agents on macOS, providing resumable file upload capabilities for your Leither network.

## Overview

tus-server is a resumable file upload server that implements the tus protocol. This guide sets up tus-server as a LaunchAgent that runs when you log in to macOS, providing reliable file upload services for your Leither network.

## Prerequisites

- Node.js (version 14 or higher)
- npm (Node Package Manager)
- macOS with user account
- Basic terminal knowledge
- **Full Disk Access** permission granted to Terminal (or your preferred terminal app)

## Step 1: Grant Full Disk Access Permission

**Important**: tus-server requires Full Disk Access to run properly as a service on macOS. This permission is required for Node.js to access files and directories across the system.

### Grant Full Disk Access to Terminal

1. Open **System Preferences** (or **System Settings** on macOS Ventura+)
2. Go to **Security & Privacy** → **Privacy** tab
3. Select **Full Disk Access** from the left sidebar
4. Click the lock icon and enter your password
5. Click the **+** button and add your Terminal application:
   - **Terminal.app**: `/Applications/Utilities/Terminal.app`
   - **iTerm2**: `/Applications/iTerm.app` (if using iTerm2)
   - **VS Code Terminal**: `/Applications/Visual Studio Code.app` (if using VS Code integrated terminal)
6. Make sure the checkbox next to your terminal app is checked
7. Restart your terminal application

### Why Full Disk Access is Required

- tus-server needs to access files across the system for file upload operations
- Node.js LaunchAgents have restricted access without Full Disk Access permission
- This permission allows tus-server to create and manage upload files in various directories
- Without this permission, tus-server may fail to start or function properly

## Step 2: Install Node.js and npm

If you don't have Node.js installed:

```bash
# Install Node.js using Homebrew (recommended)
brew install node

# Or download from https://nodejs.org/
# Verify installation
node --version
npm --version
```

## Step 3: Create tus-server Directory Structure

Create the necessary directory structure for tus-server:

```bash
# Create the main directory
sudo mkdir -p /usr/local/tus-server

# Create subdirectories
sudo mkdir -p /usr/local/tus-server/{bin,etc,logs,uploads}

# Set proper ownership
sudo chown -R $(whoami):staff /usr/local/tus-server

# Set permissions
chmod -R 755 /usr/local/tus-server
```

## Step 4: Install tus-server

Install tus-server using npm:

```bash
# Navigate to the tus-server directory
cd /usr/local/tus-server

# Initialize npm project
npm init -y

# Install tus-server
npm install tus-server

# Install additional dependencies
npm install express cors dotenv
```

## Step 5: Create tus-server Configuration

Create a configuration file for tus-server:

```bash
nano /usr/local/tus-server/etc/tus-server.conf
```

Add the following configuration:

```json
{
  "server": {
    "host": "0.0.0.0",
    "port": 1080,
    "path": "/files/"
  },
  "datastore": {
    "type": "filesystem",
    "directory": "/usr/local/tus-server/uploads"
  },
  "cors": {
    "enabled": true,
    "origin": "*",
    "methods": ["GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"],
    "headers": ["Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Upload-Offset", "Upload-Length", "Upload-Metadata", "Tus-Resumable", "Upload-Defer-Length", "Upload-Concat"]
  },
  "maxSize": "100MB",
  "timeout": 30000,
  "retry": {
    "maxRetries": 3,
    "retryDelay": 1000
  }
}
```

## Step 6: Create tus-server Startup Script

Create a startup script for tus-server:

```bash
nano /usr/local/tus-server/bin/start-tus-server.sh
```

Add the following content:

```bash
#!/bin/bash

# tus-server startup script
# This script starts the tus-server with proper configuration

# Set environment variables
export NODE_ENV=production
export TUS_SERVER_CONFIG=/usr/local/tus-server/etc/tus-server.conf

# Change to the tus-server directory
cd /usr/local/tus-server

# Start tus-server
node -e "
const tus = require('tus-server');
const express = require('express');
const cors = require('cors');
const fs = require('fs');
const path = require('path');

// Load configuration
const config = JSON.parse(fs.readFileSync(process.env.TUS_SERVER_CONFIG, 'utf8'));

// Create Express app
const app = express();

// Enable CORS
app.use(cors(config.cors));

// Create tus server
const tusServer = new tus.Server({
  path: config.server.path,
  datastore: new tus.FileStore({
    directory: config.datastore.directory
  }),
  maxSize: config.maxSize,
  timeout: config.timeout,
  retry: config.retry
});

// Mount tus server
app.all('*', tusServer.handle.bind(tusServer));

// Start server
const server = app.listen(config.server.port, config.server.host, () => {
  console.log(\`tus-server running on http://\${config.server.host}:\${config.server.port}\${config.server.path}\`);
  console.log(\`Upload directory: \${config.datastore.directory}\`);
});

// Graceful shutdown
process.on('SIGTERM', () => {
  console.log('Received SIGTERM, shutting down gracefully');
  server.close(() => {
    console.log('tus-server stopped');
    process.exit(0);
  });
});

process.on('SIGINT', () => {
  console.log('Received SIGINT, shutting down gracefully');
  server.close(() => {
    console.log('tus-server stopped');
    process.exit(0);
  });
});
"
```

Make the script executable:

```bash
chmod +x /usr/local/tus-server/bin/start-tus-server.sh
```

## Step 7: Test tus-server

Before creating the service, verify tus-server works:

```bash
# Navigate to the directory
cd /usr/local/tus-server

# Run tus-server manually
./bin/start-tus-server.sh

# If it starts successfully, stop it with Ctrl+C
```

## Step 8: Create the LaunchAgent Plist File

Create the plist configuration file:

```bash
nano ~/Library/LaunchAgents/com.tus-server.plist
```

Add the following content:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.tus-server</string>
    
    <key>ProgramArguments</key>
    <array>
        <string>/opt/homebrew/bin/node</string>
        <string>/usr/local/tus-server/app.js</string>
    </array>
    
    <key>WorkingDirectory</key>
    <string>/usr/local/tus-server</string>
    
    <key>StandardOutPath</key>
    <string>/tmp/tus-server.log</string>
    
    <key>StandardErrorPath</key>
    <string>/tmp/tus-server.err.log</string>
    
    <key>RunAtLoad</key>
    <true/>
    
    <key>KeepAlive</key>
    <true/>
    
    <key>ProcessType</key>
    <string>Background</string>
    
    <key>EnvironmentVariables</key>
    <dict>
        <key>NODE_ENV</key>
        <string>production</string>
        <key>PATH</key>
        <string>/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin</string>
    </dict>
</dict>
</plist>
```

### Plist Configuration Explained

- **Label**: Unique identifier for the service (`com.tus-server`)
- **ProgramArguments**: Direct path to Node.js and tus-server app.js file
- **WorkingDirectory**: Directory where tus-server runs (`/usr/local/tus-server`)
- **RunAtLoad**: Start service when loaded (at login)
- **KeepAlive**: Always keep the service running (restart if it crashes)
- **ProcessType**: Run as background process
- **StandardOutPath**: Location for standard output logs (`/tmp/tus-server.log`)
- **StandardErrorPath**: Location for error logs (`/tmp/tus-server.err.log`)
- **EnvironmentVariables**: Environment variables including NODE_ENV and PATH

## Step 9: Set Proper Permissions

```bash
chmod 644 ~/Library/LaunchAgents/com.tus-server.plist
```

## Step 10: Validate the Plist File

```bash
plutil -lint ~/Library/LaunchAgents/com.tus-server.plist
```

## Step 11: Load the Service

```bash
# Load the service
launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.tus-server.plist

# Check if it's running
launchctl list | grep tus-server
```

Expected output:
```
3149    0    com.tus-server
```

## Step 12: Check Logs

```bash
# View standard output
cat /tmp/tus-server.log

# View errors
cat /tmp/tus-server.err.log

# Monitor logs in real-time
tail -f /tmp/tus-server.log
```

## Step 13: Test After Reboot

```bash
sudo reboot
```

After logging back in, wait about 60 seconds, then check:

```bash
launchctl list | grep tus-server
cat /tmp/tus-server.log
```

## Managing the Service

### Unload (Stop and Remove)
```bash
launchctl bootout gui/$(id -u)/com.tus-server
```

### Reload After Changes
```bash
launchctl bootout gui/$(id -u)/com.tus-server
launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.tus-server.plist
```

## Testing tus-server

### Basic Upload Test
```bash
# Create a test file
echo "Hello, tus-server!" > /tmp/test.txt

# Upload using curl
curl -X POST http://localhost:1080/files/ \
  -H "Upload-Length: 18" \
  -H "Upload-Metadata: filename dGVzdC50eHQ=" \
  --data-binary @/tmp/test.txt
```

### Resume Upload Test
```bash
# Start an upload
curl -X POST http://localhost:1080/files/ \
  -H "Upload-Length: 18" \
  -H "Upload-Metadata: filename dGVzdC50eHQ=" \
  --data-binary @/tmp/test.txt

# Resume the upload (replace with actual upload URL)
curl -X PATCH http://localhost:1080/files/[upload-id] \
  -H "Upload-Offset: 0" \
  -H "Tus-Resumable: 1.0.0" \
  --data-binary @/tmp/test.txt
```

## Integration with Leither

### Configure Leither to Use tus-server

Add tus-server configuration to your Leither setup:

```bash
# Edit Leither configuration
nano /usr/local/darwin/Systemvar.json
```

Add tus-server endpoint:
```json
{
  "tus_server": {
    "endpoint": "http://localhost:1080/files/",
    "max_size": "100MB",
    "timeout": 30000
  }
}
```

### Upload Files via tus-server

```bash
# Upload a file using tus protocol
curl -X POST http://localhost:1080/files/ \
  -H "Upload-Length: $(stat -f%z /path/to/file)" \
  -H "Upload-Metadata: filename $(base64 -i /path/to/file)" \
  --data-binary @/path/to/file
```

## Troubleshooting

### Service Shows "-" Instead of PID
Check logs: `cat /tmp/tus-server.err.log`

### "Input/output error" When Loading
1. Service is already loaded (unload first)
2. Invalid plist syntax (run `plutil -lint`)

### "Operation not permitted" Errors
Executable is in protected location. Ensure proper permissions on `/usr/local/tus-server`.

### "Full Disk Access" Permission Denied
If tus-server fails to start due to permission issues:

1. **Check Full Disk Access Status**:
   ```bash
   # Check if Terminal has Full Disk Access
   sqlite3 ~/Library/Application\ Support/com.apple.TCC/TCC.db "SELECT * FROM access WHERE service='kTCCServiceSystemPolicyAllFiles';"
   ```

2. **Grant Full Disk Access** (if not already done):
   - Go to **System Preferences** → **Security & Privacy** → **Privacy** → **Full Disk Access**
   - Add your Terminal application
   - Restart Terminal and try again

3. **Alternative: Use LaunchDaemon instead of LaunchAgent**:
   - Move plist to `/Library/LaunchDaemons/` (requires sudo)
   - This gives the service system-level permissions

### Service Doesn't Start After Reboot
Increase `StartInterval` to 90 or 120 seconds

### tus-server Not Responding
1. Check if port 1080 is available: `lsof -i :1080`
2. Verify Node.js is installed: `node --version`
3. Check tus-server logs: `tail -f /tmp/tus-server.log`

### Upload Failures
1. Check upload directory permissions: `ls -la /usr/local/tus-server/uploads`
2. Verify disk space: `df -h /usr/local/tus-server`
3. Check tus-server configuration: `cat /usr/local/tus-server/etc/tus-server.conf`

## Important Notes

1. **Don't use sudo** to load LaunchAgents
2. **Use absolute paths** in the plist file
3. **The 60-second delay** prevents failures during boot
4. **tus-server requires Node.js** - ensure it's installed and accessible
5. **Upload directory** must be writable by the user running the service

## File Locations

- **Plist**: `~/Library/LaunchAgents/com.tus-server.plist`
- **Binary**: `/usr/local/tus-server/app.js`
- **Configuration**: `/usr/local/tus-server/etc/tus-server.conf`
- **Logs**: `/tmp/tus-server.log` and `/tmp/tus-server.err.log`
- **Uploads**: `/usr/local/tus-server/uploads/`

## Security Considerations

1. **Firewall**: Configure macOS firewall to allow port 1080
2. **CORS**: Adjust CORS settings in configuration for production
3. **Authentication**: Consider adding authentication for production use
4. **SSL/TLS**: Use HTTPS in production environments
5. **File Size Limits**: Configure appropriate maxSize limits

---

**Setup completed successfully! tus-server will now start automatically when you log in to macOS and provide resumable file upload capabilities for your Leither network.**
