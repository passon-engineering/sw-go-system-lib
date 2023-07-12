package system

import (
	"bufio"
	"fmt"
	"os/exec"
)

// Runs a command in a shell and captures its output.
//
// It executes the specified command using the shell, reads its stdout pipe,
// and returns the output as a byte slice. Any error that occurred during
// command execution or output retrieval is returned.
//
// Parameters:
//   - command: string - the command to be executed
//
// Returns:
//   - []byte: the output of the command
//   - error: an error if any occurred during command execution or output retrieval
//
// Example usage:
//
//	output, err := RunCommandGetOutput("ls -l")
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

// Checks if a command exists on the system.
//
// It attempts to execute the specified command with the "-v" flag, which outputs
// the command's version information. If the command is found and executable,
// it returns true. Otherwise, it returns false.
//
// Parameters:
//   - command: string - the command to check
//
// Returns:
//   - bool: true if the command exists, false otherwise
//   - error: an error if any occurred during command execution
//
// Example usage:
//
//	exists, err := CommandExists("ls")
//	if err != nil {
//	  // handle error
//	}
//	if exists {
//	  fmt.Println("ls command exists!")
//	} else {
//	  fmt.Println("ls command does not exist!")
//	}
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
