package networking

import (
	"errors"
	"net"
	"time"
)

/*
GetNetworkExternalIP retrieves the external IP address of the machine.

It iterates through the network interfaces, excluding loopback interfaces
and those that are down. For each valid interface, it retrieves the IP addresses
and returns the first non-loopback IPv4 address found.

The function makes multiple attempts with a retry delay in case the IP address
cannot be retrieved immediately. If the maximum number of attempts is reached
without success, it returns an error.

Returns:
  - string: the external IP address
  - error: an error if the IP address retrieval fails

Example usage:
  ip, err := GetNetworkExternalIP()
*/
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



/*
IsValidIPv4 checks if the given IP address is a valid IPv4 address.

It uses the net.ParseIP function to parse the IP address string. If the parsed
IP address is nil, it indicates that the IP address is invalid and the function
returns false. If the parsed IP address is not nil but does not represent an
IPv4 address, the function also returns false. Otherwise, it considers the
IP address as a valid IPv4 address and returns true.

Parameters:
  - ip: string - the IP address to validate

Returns:
  - bool: true if the IP address is a valid IPv4 address, false otherwise

Example usage:
  isValid := IsValidIPv4("192.168.0.1")
*/
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