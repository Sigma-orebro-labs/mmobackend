package main

import (
	"log"
	"net"
	"time"
)

const (
	laddr = "127.0.0.1:3333"
	maxNumConnections = 100
)

func main() {
	numConnections := 0
	for {
		if numConnections < maxNumConnections {
			go connect()
			numConnections++
		}
	}
}

func connect() {
	conn, err := net.Dial("tcp", laddr)
	
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		write(conn)
		message, err := receive(conn)
		
		if err != nil {
			log.Fatal(err)
		}
		
		log.Println("Command: ", message[2])
		
		time.Sleep(100 * time.Millisecond)
	}
}

func write(conn net.Conn) error {
	message := make([]byte, 7)
	
	message[0] = 0xF0;
	message[1] = 3;
	message[2] = 1;
	message[3] = 101; 	// x
	message[4] = 102; 	// y
	message[5] = 103; 	// z
	message[6] = 0xCC;
	
	_, err := conn.Write(message)
	
	return err
}

func receive(conn net.Conn) ([]byte, error) {
	message := make([]byte, 7)
	_, err := conn.Read(message)
	
	return message, err
}