package system

import (
	"bufio"
	"fmt"
	"os/exec"
)

// RunCommandGetOutput executes the specified command in a shell and returns its output.
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
