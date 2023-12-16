// main.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var ipStr string

	// Check if command-line arguments are provided
	if len(os.Args) > 1 {
		ipStr = os.Args[1]
	} else {
		// Read from standard input (stdin) if no command-line arguments are provided
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			ipStr = scanner.Text()
		} else {
			fmt.Println("No input provided.")
			os.Exit(1)
		}
	}

	ip := net.ParseIP(strings.TrimSpace(ipStr))
	if ip == nil {
		fmt.Println("Invalid IP address:", ipStr)
		os.Exit(1)
	}

	for i := 1; i <= 255; i++ {
		ip := net.IPv4(ip[12], ip[13], ip[14], byte(i))
		fmt.Printf("%s/%d\n", ip.String(), i)
	}
}
