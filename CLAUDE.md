# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Leither is a decentralized cloud operating system that provides an ultra-lightweight platform for building decentralized applications. The system uses "Mimei" (information containers) as its core concept to store both data and applications, enabling decentralized network functionality.

## Architecture

### Core Components

- **Mimei System**: Information containers that can store files or databases
- **Application System**: Applications that run within the Leither ecosystem
- **Network Layer**: DHT-based decentralized networking compatible with IPFS
- **API Layer**: RPC-based interface using Hprose protocol

### Key Directories

- `lapi/`: Core API package providing RPC functionality
- `example/`: Example applications and usage patterns
- `www/`: Web interface and documentation
- `api/`: API documentation
- `doc/`: System documentation

## Development Commands

### Building and Running

```bash
# Initialize a Leither node
./Leither init -p 4800 -b mimei.org

# Run as background service
./Leither run -d

# Stop the service
./Leither stop
```

### Go Development

```bash
# Build the project
go build

# Run tests
go test ./...

# Install dependencies
go mod tidy
```

### Example Applications

```bash
# Run the hello world example
cd example/go/helloleither
go run main.go

# Run RPC example
cd example/rpc
go run example.go
```

## API Architecture

### Core Interfaces

- `LApi`: Main interface for node functionality including:
  - `IBackEnd`: Backend operations (session management, logging)
  - `IAuth`: Authentication and authorization
  - `IVarAct`: Variable operations and actions
  - `IMiMei`: Mimei container operations
  - `INet`: Network operations

- `IRPC`: Remote procedure call interface for cross-node communication

### Key API Patterns

1. **RPC Access**: Use `InitLApiStubByUrl("127.0.0.1:4800")` to create RPC stubs
2. **Backend Access**: Use `GetLApi()` for in-node application execution
3. **Session Management**: Use `CreateSession()` for state management
4. **Mimei Operations**: Create, open, and manipulate Mimei containers

## Mimei System

Mimei containers are the fundamental building blocks:

- **File Mimei**: Store and version files
- **Database Mimei**: Redis-compatible database operations
- **Application Mimei**: Store and run applications

Key operations:
- `MMCreate()`: Create new Mimei
- `MMOpen()`: Open existing Mimei
- `MMBackup()`: Create version snapshots
- `MMSync()`: Synchronize across nodes

## Network Features

- **DHT Network**: Decentralized hash table for peer discovery
- **IPFS Compatibility**: Full compatibility with IPFS protocol
- **Load Balancing**: Browser-based node selection for optimal performance
- **Fault Tolerance**: Multiple provider support for data redundancy

## Development Workflow

1. **Create Application**: Build applications using Leither API
2. **Package as Mimei**: Upload application to Mimei container
3. **Publish**: Distribute application across the network
4. **Synchronize**: Nodes sync applications and data automatically

## Testing

The project includes example applications in the `example/` directory that demonstrate:
- Basic RPC communication
- Mimei operations
- Application execution patterns

## Documentation

- `README.md`: Chinese documentation with comprehensive usage examples
- `README_EN.md`: English documentation
- `api/`: Detailed API documentation
- `doc/`: System architecture and concept documentation