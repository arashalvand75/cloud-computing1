package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	// Check if the memory limit argument is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: container-runtime <hostname> <memory_limit_in_mb>")
		return
	}

	hostname := os.Args[1]
	memoryLimitMB, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid memory limit: %v\n", err)
		return
	}

	// Create a cgroup for memory limitation
	memCgroup, err := createMemoryCgroup(memoryLimitMB)
	if err != nil {
		fmt.Printf("Failed to create memory cgroup: %v\n", err)
		return
	}
	defer cleanupCgroup(memCgroup)

	// Join the cgroup before running the container process
	err = joinCgroup(memCgroup)
	if err != nil {
		fmt.Printf("Failed to join cgroup: %v\n", err)
		return
	}

	// Run the container process (bash in this example)
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Container process exited with error: %v\n", err)
	}
}

func createMemoryCgroup(memoryLimitMB int) (string, error) {
	// Create a new cgroup for memory limitation
	memCgroup := "/sys/fs/cgroup/memory/container"
	err := os.Mkdir(memCgroup, 0755)
	if err != nil {
		return "", err
	}

	// Set the memory limit for the cgroup
	memoryLimitBytes := memoryLimitMB * 1024 * 1024
	err = os.WriteFile(memCgroup+"/memory.limit_in_bytes", []byte(strconv.Itoa(memoryLimitBytes)), 0644)
	if err != nil {
		return "", err
	}

	return memCgroup, nil
}

func cleanupCgroup(memCgroup string) {
	// Remove the cgroup directory
	os.RemoveAll(memCgroup)
}

func joinCgroup(memCgroup string) error {
	// Join the cgroup by writing the current process ID to the tasks file
	pid := os.Getpid()
	err := os.WriteFile(memCgroup+"/tasks", []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return err
	}

	// Set the appropriate namespace for the current process
	err = syscall.Unshare(syscall.CLONE_NEWNS)
	if err != nil {
		return err
	}

	return nil
}