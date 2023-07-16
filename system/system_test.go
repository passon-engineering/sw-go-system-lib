package system

import (
	"fmt"
	"strings"
	"testing"
)

func TestFunctions(t *testing.T) {
	testRunCommandGetOutput(t)
	testRunCommand(t)
	testRunCommandExist(t)
}

func testRunCommandGetOutput(t *testing.T) {
	output, err := RunCommandGetOutput("echo 'Hello, World!'")
	if err != nil {
		t.Errorf("Failed to run command: %v", err)
	}

	expected := []byte("Hello, World!")
	if string(output) != string(expected) {
		t.Errorf("Unexpected command output. Expected: %s, Got: %s", expected, output)
	}

	// Print the command output for visual inspection
	t.Logf("Command output: %s", output)
}

func testRunCommand(t *testing.T) {
	output, err := RunCommand("echo 'Hello, World!'")
	if err != nil {
		t.Errorf("Failed to run command: %v", err)
	}

	// Trim any trailing newline or space
	output = strings.TrimSpace(output)

	expected := "Hello, World!"
	if output != expected {
		t.Errorf("Unexpected command output. Expected: %s, Got: %s", expected, output)
	}

	// Print the command output for visual inspection
	t.Logf("Command output: %s", output)
}

func testRunCommandExist(t *testing.T) {
	command := "ls"
	output, err := CommandExists(command)
	if err != nil {
		fmt.Printf("Error occurred while checking command existence: %v\n", err)
	}

	expected := true
	if output != expected {
		t.Errorf("Unexpected command output. Expected: %v, Got: %v", expected, output)
	}

	command = "fake-command"
	output, err = CommandExists(command)
	if err != nil {
		fmt.Printf("Error occurred while checking command existence: %v\n", err)
	}

	expected = false
	if output != expected {
		t.Errorf("Unexpected command output. Expected: %v, Got: %v", expected, output)
	}
}
