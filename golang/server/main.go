package main

import (
	"log"
	"net"
	"github.com/Sigma-orebro-labs/mmobackend/golang/server/handlers"
)

const (
	ip = "127.0.0.1"
	tcpPort = 11000
	udpPort = 11001
)

func main() {
	block := make(chan bool)
	
	go listenTcp()
	go listenUdp()
	
	<-block
}

func listenTcp() {
	addr := net.TCPAddr {
		Port: tcpPort,
		IP: net.ParseIP(ip),
	}
	
	l, err := net.ListenTCP("tcp", &addr)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer l.Close()
	
	log.Println("Tcp: ", ip, ":", tcpPort)
	
	for {
		// This will block until we recieve a connection
		conn, err := l.Accept()
		
		if err != nil {
			log.Println(err)
			continue
		}
		
		go handlers.HandleTcp(conn)
	}
}

func listenUdp() {
	addr := net.UDPAddr {
		Port: udpPort,
		IP: net.ParseIP(ip),
	}
	
	buffer := make([]byte, 1024)
	
	conn, err := net.ListenUDP("udp", &addr)
	
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Udp: ", ip, ":", udpPort)
	
	for {
		_, _, err = conn.ReadFromUDP(buffer)
		
		if err != nil {
			log.Println(err)
			continue
		}
	}
}