Network and Container Projects
This repository contains three distinct problems related to networking, containerization, and server implementation. Each project focuses on different aspects of system administration and software development.
problem 1: Network Namespace Topology
Explores various network namespace topologies using Linux network namespaces, bridges, and virtual Ethernet interfaces.
Key Features:

Single-server with Router topology
Single-server without Router topology
Multi-server topology
Uses Linux network namespaces, bridges, and virtual Ethernet interfaces

Requirements:

Linux operating system
Root or sudo access
iproute2 package

problem 2: Simple Container Runtime
A basic container runtime implementation in Go, creating isolated environments with namespace separation.
Key Features:

Creates new namespaces for containers
Sets up a new root filesystem
Optional memory usage limiting
Written in Go

Requirements:

Go programming language
Ubuntu 20.04 root filesystem

problem 3: Status Server
A simple HTTP server with Docker containerization, providing a single endpoint for status management.
Key Features:

Single endpoint: /api/v1/status
Supports GET and POST HTTP methods
Dockerized for easy deployment

Requirements:

Python 
Docker

Usage
Each project has its own specific usage instructions. Please refer to the individual README files within each project directory for detailed information on building, running, and using the respective projects.
Contributing
Contributions to any of the projects are welcome. Please open an issue or submit a pull request in the respective project directory.
License
All projects in this repository are licensed under the MIT License, unless otherwise specified in the individual project directories.
