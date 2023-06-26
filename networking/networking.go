package networking

import (
	"errors"
	"net"
	"time"
)

// GetNetworkExternalIP attempts to retrieve the external IP address.
// It iterates through network interfaces and returns the first non-loopback IPv4 address found.
// It tries for a maximum of 5 attempts with a 5-second delay between each attempt.
func GetNetworkExternalIP() (string, error) {
	const maxAttempts = 5
	const retryDelay = 5 * time.Second

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		ifaces, err := net.Interfaces()
		if err != nil {
			return "", err
		}

		for _, iface := range ifaces {
			if iface.Flags&net.FlagUp == 0 {
				continue // interface is down
			}
			if iface.Flags&net.FlagLoopback != 0 {
				continue // loopback interface
			}

			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				if ip == nil || ip.IsLoopback() {
					continue
				}

				ip = ip.To4()
				if ip == nil {
					continue // not an IPv4 address
				}

				return ip.String(), nil
			}
		}

		time.Sleep(retryDelay)
	}

	return "", errors.New("failed to retrieve external IP address")
}



// isValidIPv4 checks whether the given string is a valid IPv4 address.
func IsValidIPv4(ip string) bool {
	if parsedIP := net.ParseIP(ip); parsedIP == nil {
		// Invalid IPv4 address
		return false
	} else if parsedIP.To4() == nil {
		// Not an IPv4 address
		return false
	}

	return true
}