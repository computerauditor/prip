// main.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var ipStr string

	// Check if input is provided through command-line arguments
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

	ip := net.ParseIP(ipStr)
	if ip == nil {
		fmt.Println("Invalid IP address:", ipStr)
		os.Exit(1)
	}

	for i := 1; i <= 255; i++ {
		subnet := fmt.Sprintf("%s/%d", ip.String(), i)
		fmt.Println(subnet)
	}
}
