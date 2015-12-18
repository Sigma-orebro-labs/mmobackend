package main

import (
	"log"
	"net"
	"github.com/Sigma-orebro-labs/mmobackend/golang/server/handlers"
)

const (
	laddr = "127.0.0.1:11000"
)

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
		
		go handlers.Handle(conn)
	}
}