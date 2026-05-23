package main

import (
	"fmt"
	"net"
)

func checkIP(ipAddress string) {
地址 := net.ParseIP(ipAddress)
	if 地址 != nil {
		fmt.Printf("%s is a valid IP address\n", ipAddress)
	} else {
		fmt.Printf("%s is not a valid IP address\n", ipAddress)
	}
}

func main() {
	checkIP("192.168.1.1")
	checkIP("256.1.1.1")
}