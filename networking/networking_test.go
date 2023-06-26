package networking

import (
	"testing"
)

func TestFunctions(t *testing.T) {
	testGetNetworkExternalIP(t)
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