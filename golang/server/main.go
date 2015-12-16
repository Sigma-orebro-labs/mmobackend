package main

import (
	"log"
	"net"
	"io"
)

const (
	laddr = "127.0.0.1:11000"
	
	maxMessageBodyLength byte = 128;
	messageHeaderMarker byte = 0xF0 	// 1111 0000
	messageFooterMarker byte = 0xCC 	// 1100 1100
	messageHeaderLength byte = 3
	messageFooterLength byte = 1
	
	// Commands
	getCurrentUserPositionCommand byte = 1
    getCurrentUserPositionResponse byte = 2
)

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
		
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	
	response := make([]byte, maxMessageBodyLength)
	response[0] = messageHeaderMarker;
	responseBaseSize := messageHeaderLength + messageFooterLength
	
	header := make([]byte, messageHeaderLength)
	body := make([]byte, maxMessageBodyLength)
	footer := make([]byte, messageFooterLength)
	
	for {
		// Read will read len(b) number of bytes
		_, err := conn.Read(header)
		
		if err != nil {
			if err == io.EOF {
				break	
			}
			
			log.Println(err)
			break
		}
		
		if header[0] != messageHeaderMarker {
			log.Println("Invalid message header")
			break 
		}
		
		commandCode := header[1];
		bodyLength := header[2];
		
		_, err = conn.Read(body[0:bodyLength])
		
		if err != nil {
			log.Println(err)
		    log.Println("Something went wrong while reading the message body")
			break
		}
		
		_, err = conn.Read(footer)
	
		if err != nil {
			log.Println(err)
			break
		}
		
		if footer[0] != messageFooterMarker {
			log.Fatal("Invalid message footer")
			break
		}
		
		switch (commandCode) {
		case getCurrentUserPositionCommand:
			
			response[1] = getCurrentUserPositionResponse;
			response[2] = 3; 	// 3 byte response body (x, y, z coordinates)
		    response[3] = 101; 	// x
		    response[4] = 102; 	// y
		    response[5] = 103; 	// z
		    response[6] = messageFooterMarker;
			
			_, err = conn.Write(response[0:responseBaseSize + 3])	
			if err != nil {
		 		log.Println(err)
		 	}
				
		    break;		
		default:
		    break;
		}
	}
}