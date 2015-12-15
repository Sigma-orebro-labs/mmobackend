package main

import (
	"log"
	"net"
	"time"
)

const (
	maxNumConnections = 1000
)

func main() {
	numConnections := 0
	for {
		if numConnections <= maxNumConnections {
			go connect()
			numConnections++
		}
	}
}

func connect() {
	conn, err := net.Dial("tcp", "localhost:3333")
	
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		write(conn)
		time.Sleep(100 * time.Millisecond)
	}
}

func write(conn net.Conn) error {
	message := "Hello from client \n"
	_, err := conn.Write([]byte(message))
	
	return err
}