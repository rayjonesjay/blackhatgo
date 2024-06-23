package tcp

import (
	"fmt"
	"net"
)

func FastScanner() {

	for portNo := 1 ; portNo <= 1024; portNo++{
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d",j)
			conn , err := net.Dial("tcp", address)
			if err != nil {
				return 
			}
			conn.Close()
			fmt.Printf("[+] %d open\n",j)
		}(portNo)
	}
}