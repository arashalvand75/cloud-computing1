# Simple Container Runtime

This is a simple container runtime implementation in Go. It creates a new namespace for the container, sets up a new root filesystem, and optionally limits the memory usage.

## Prerequisites

- Go programming language installed
- Ubuntu 20.04 root filesystem (e.g., extracted from the `ubuntu:20.04` Docker image)

## Building

1. Clone the repository or copy the source code.
2. Build the binary:

## Running

1. Make sure the `containerRootPath` variable in the source code points to the correct location of the Ubuntu 20.04 root filesystem.
2. Run the binary with the desired hostname and optional memory limit (in megabytes):

3. After running the CLI, you'll be dropped into a new bash shell inside the container namespace with the specified hostname.
4. Run `ps fax` to verify that the `bash` process is running as PID 1.

## Notes

- This is a simplified implementation and may not include all the features and security considerations of a production-ready container runtime.
- The container root filesystem is not persistent and will be removed after exiting the container.
- The implementation assumes that the Ubuntu 20.04 root filesystem is available and accessible at the specified `containerRootPath`.




