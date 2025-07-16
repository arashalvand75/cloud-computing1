# Network Namespace Topology

This project explores different network namespace topologies and routing solutions using Linux network namespaces, bridges, and virtual Ethernet interfaces.

## Table of Contents

1. [Introduction](#introduction)
2. [Topology Scenarios](#topology-scenarios)
3. [Solutions](#solutions)
4. [Requirements](#requirements)
5. [Usage](#usage)
6. [Contributing](#contributing)
7. [License](#license)

## Introduction

This project demonstrates various network topologies using Linux network namespaces. It covers scenarios ranging from a single-server setup with a router to multi-server configurations without a dedicated router.

## Topology Scenarios

1. **Single-server with Router**: Network namespaces connected via bridges and a router namespace.
2. **Single-server without Router**: Network namespaces connected via bridges, with routing handled by the root namespace.
3. **Multi-server**: Network namespaces distributed across multiple servers, connected at Layer 2.

## Solutions

### Single-server with Router

- Uses network namespaces for nodes and router
- Connects namespaces using virtual Ethernet interfaces and bridges
- Implements routing between subnets using the router namespace

### Single-server without Router

- Uses network namespaces for nodes
- Connects namespaces using virtual Ethernet interfaces and bridges
- Implements routing between subnets using IP forwarding and policy routing in the root namespace

### Multi-server

Two approaches are discussed:

1. **Using Bridged Networking**:
   - Connects servers at Layer 2 using a physical or virtual switch
   - Simpler to implement but has limitations in scalability and isolation

## Requirements

- Linux operating system
- Root or sudo access
- `iproute2` package

## Usage

Detailed implementation scripts are not provided in this repository. However, the README offers explanations and conceptual solutions for each topology scenario.

To implement these topologies, you would need to:

1. Create network namespaces
2. Set up virtual Ethernet interfaces and bridges
3. Configure IP addresses and routing rules
4. Enable IP forwarding where necessary

For the multi-server scenario, additional network configuration between servers would be required.

## Contributing

Contributions to improve the explanations, add implementation scripts, or explore additional scenarios are welcome. Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
