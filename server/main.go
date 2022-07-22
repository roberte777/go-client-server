package main

import (
	"fmt"
	"net"
	"os"
)

//make a tcp server that listens on port 8080
func main() {
	const (
		PORT = "80"
		HOST = "172.16.0.6"
		TYPE = "tcp"
	)
	//listen on port 8080
	l, err := net.Listen(TYPE, HOST+":"+PORT)
	fmt.Println("Listening on port " + PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//run loop forever (or until ctrl-c)
	for {
		//accept a connection
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//handle the connection
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println("Handling connection")
	buffer := make([]byte, 5)
	for {
		n, err := c.Read(buffer)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c.Write([]byte(string(buffer[:n])))

	}
}
