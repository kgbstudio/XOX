package main

import (
	"fmt"
	"net"
)

// Function to check if an IP address is valid
func checkIP(ipAddress string) {
	parsedIP := net.ParseIP(ipAddress)
	if parsedIP != nil {
		fmt.Printf("%s is a valid IP address\n", ipAddress)
	} else {
		fmt.Printf("%s is not a valid IP address\n", ipAddress)
	}
}

func main() {
	// Test cases
	ipAddresses := []string{
		"192.168.1.1",
		"256.1.1.1",
		"::1", // IPv6
		"2001:0db8:85a3:0000:0000:8a2e:0370:7334", // IPv6
		"2001:0db8:85a3:0000:0000:8a2e:0370:733g", // Invalid IPv6
	}

	for _, ipAddress := range ipAddresses {
		checkIP(ipAddress)
	}
}