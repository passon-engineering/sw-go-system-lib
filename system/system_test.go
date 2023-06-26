package system

import (
	"testing"
)

func TestFunctions(t *testing.T) {
	testRunCommandGetOutput(t)
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