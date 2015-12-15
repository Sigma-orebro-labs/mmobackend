package main

import (
	"log"
	"net"
)

const (
	laddr = "127.0.0.1:3333"
	
	maxMessageBodyLength byte = 128;
	messageHeaderMarker byte = 0xF0 	// 1111 0000
	messageFooterMarker byte = 0xCC 	// 1100 1100
	messageHeaderLength byte = 3
	messageFooterLength byte = 1
	
	// Commands
	getCurrentUserPositionCommand byte = 1
    getCurrentUserPositionResponse byte = 2
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
	l, err := net.Listen("tcp", laddr)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer l.Close()
	
	log.Println(laddr)
	
	for {
		// This will block until we recieve a connection
		conn, err := l.Accept()
		
		if err != nil {
			log.Println(err)
			continue
		}
		
		log.Println("Connection found")
		
		go handle(conn)
		
		numConnections++
		log.Println("Number of connections:", numConnections)	
	}
}

func handle(conn net.Conn) {
	response := make([]byte, maxMessageBodyLength)
	response[0] = messageHeaderMarker;
	responseBaseSize := messageHeaderLength + messageFooterLength
	
	header := make([]byte, messageHeaderLength)
	footer := make([]byte, messageFooterLength)
	
	// Read will read len(b) number of bytes
	_, err := conn.Read(header)
	
	if err != nil || header[0] != messageHeaderMarker {
		log.Println(err)
		log.Fatal("Something went wrong while reading the message header")
	}
	
	bodyLength := header[1];
	commandCode := header[2];
	
	body := make([]byte, bodyLength)
	_, err = conn.Read(body)
	
	if err != nil {
		log.Println(err)
	    log.Fatal("Something went wrong while reading the message body")
	}
	
	_, err = conn.Read(footer)

	if err != nil || footer[0] != messageFooterMarker {
		log.Println(err)
	    log.Fatal("Something went wrong while reading the message footer")
	}
	
	switch (commandCode) {
	case getCurrentUserPositionCommand:
		
		response[1] = 3; 	// 3 byte response body (x, y, z coordinates)
	    response[2] = getCurrentUserPositionResponse;
	    response[3] = 101; 	// x
	    response[4] = 102; 	// y
	    response[5] = 103; 	// z
	    response[6] = messageFooterMarker;
		
		_, err = conn.Write(response[0:responseBaseSize + 3])	
		if err != nil {
	 		log.Fatal(err)
	 	}
			
	    break;		
	default:
	    break;
	}
}