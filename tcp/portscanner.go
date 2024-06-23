package tcp

import (
	"fmt"
	"net"
)

func Hello() {
	fmt.Println("1234")
}

func Scan() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")

	if err == nil {
		fmt.Println("connection success")
	} else {
		fmt.Println("failed")
	}
}
