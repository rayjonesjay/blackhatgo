package tcp

import (
	"fmt"
	"net"
)

func ScanSlow() {
	for portNumber := 20; portNumber <= 60; portNumber++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", portNumber)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// this means connection is closed or filtered
			continue
		}
		// if a connection was a success close the connection
		conn.Close()
		fmt.Printf("%d open\n", portNumber)
	}
}
