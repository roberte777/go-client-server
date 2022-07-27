package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

//make a client that connects to tcp server
func main() {
	const CONNECTION = "localhost:80"
	//connect to server
	c, err := net.Dial("tcp", CONNECTION)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	//make a buffer to read data into
	//read data from server

	go read_data(c)
	go write_data(c)
	//for loop to keep main running
	for {
		//do nothing
		time.Sleep(time.Second * 1)
	}
}
func read_data(c net.Conn) {
	//make a buffer to read data into
	buf := make([]byte, 1024)
	for {
		//read data from server
		n, err := c.Read(buf)
		if err != nil {
			break
		}
		//print data to screen
		fmt.Println("FROM SERVER: " + string(buf[:n]))

	}

}

//write user input
func write_data(c net.Conn) {
	for {
		//get user input
		var input string
		reader := bufio.NewScanner(os.Stdin)
		if reader.Scan() {
			input = reader.Text()
		}

		//send user input to server
		c.Write([]byte(input))

	}
}
