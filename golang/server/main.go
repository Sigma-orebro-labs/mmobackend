package main

import (
	"log"
	"net"
)

var numConnections = 0

type client struct {
	Conn net.Conn
}

func (c *client) Read(buffer []byte) bool {
	_, err := c.Conn.Read(buffer)
	if err != nil {
		c.Conn.Close()
		log.Println(err)
		return false
	}
	
	return true
} 

func main() {
	l, err := net.Listen("tcp", "localhost:3333")
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer l.Close()
	
	log.Println("Listening on localhost:3333")
	
	for {
		conn, err := l.Accept()
		
		if err != nil {
			log.Println(err)
		}
		
		go handle(conn)
		
		numConnections++
		log.Println("Number of connections:", numConnections)	
	}
}

func handle(conn net.Conn) {
	c := client { conn }
	buffer := make([]byte, 2048)
	
	for c.Read(buffer) {
		//log.Println("Received message: ", string(buffer))
		
		_, err := conn.Write([]byte("Hello from server \n"))
		
		if err != nil {
			log.Println(err)
		}
	}
}