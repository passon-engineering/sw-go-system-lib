package system

import (
	"bufio"
	"fmt"
	"os/exec"
)

/*
RunCommandGetOutput runs a command in a shell and captures its output.

It executes the specified command using the shell, reads its stdout pipe,
and returns the output as a byte slice. Any error that occurred during
command execution or output retrieval is returned.

Parameters:
  - command: string - the command to be executed

Returns:
  - []byte: the output of the command
  - error: an error if any occurred during command execution or output retrieval

Example usage:
  output, err := RunCommandGetOutput("ls -l")
*/
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
