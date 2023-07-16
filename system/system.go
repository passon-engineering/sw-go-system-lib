package system

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

// Executes the specified command in a shell and returns the output as a byte slice.
// Any error occurred during command execution or output retrieval is also returned.
//
// Example:
//
//	output, err := RunCommandGetOutput("ls -l")
//
// Parameters:
//
//	command string: the command to be executed
//
// Returns:
//
//	[]byte: the output of the command
//	error: an error if occurred during command execution or output retrieval
func RunCommandGetOutput(command string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", command)

	// Create a pipe for the command's output
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe for command: %w", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	scanner := bufio.NewScanner(cmdReader)
	defer cmd.Wait()

	var result []byte

	for scanner.Scan() {
		result = append(result, scanner.Bytes()...)
	}

	return result, nil
}

// Executes a specified command based on the operating system and returns the output as a string.
// On Windows, it runs the command using 'cmd', and '/bin/sh' is used on Unix-based systems.
// If an error occurs during command execution, a wrapped error with a description is returned.
//
// Parameters:
//
//	cmdLine string: the command line to be executed
//
// Returns:
//
//	string: the output of the command execution
//	error: an error if occurred during command execution
func RunCommand(cmdLine string) (string, error) {
	var cmd *exec.Cmd

	// Determine how to run the command based on the operating system
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdLine)
	} else {
		cmd = exec.Command("/bin/sh", "-c", cmdLine)
	}

	// Create buffers to capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// If an error occurred, wrap it in a more descriptive error
	if err != nil {
		return "", fmt.Errorf("command failed: %s, error: %s, stderr: %s", cmdLine, err, stderr.String())
	}

	// Return the stdout output
	return stdout.String(), nil
}

// Checks the existence of a specified command on the system.
// It tries to execute the command with '-v' flag, which normally outputs the command version.
// If the command is found and is executable, it returns true and nil error.
// If the command does not exist, it returns false and nil error.
// If an error occurs during the command check, it returns false and the occurred error.
//
// Example:
//
//	exists, err := CommandExists("ls")
//	if exists {
//	  fmt.Println("ls command exists!")
//	} else {
//	  fmt.Println("ls command does not exist!")
//	}
//
// Parameters:
//
//	command string: the command to be checked
//
// Returns:
//
//	bool: true if the command exists, false otherwise
//	error: an error if occurred during command check
func CommandExists(command string) (bool, error) {
	cmd := exec.Command("command", "-v", command)
	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			// Command not found
			return false, nil
		}
		return false, fmt.Errorf("failed to run command: %w", err)
	}
	return true, nil
}
