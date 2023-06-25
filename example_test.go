package utility

import (
	"testing"
)

func TestUtilityFunctions(t *testing.T) {
	testGetNetworkExternalIP(t)
	testRunCommandGetOutput(t)
}



func testGetNetworkExternalIP(t *testing.T) {
	ip, err := GetNetworkExternalIP()
	if err != nil {
		t.Errorf("Failed to retrieve external IP address: %v", err)
	}

	// Validate the IP address format
	if !IsValidIPv4(ip) {
		t.Errorf("Returned IP address is not valid: %s", ip)
	}

	// Print the retrieved IP address for visual inspection
	t.Logf("Retrieved external IP address: %s", ip)
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