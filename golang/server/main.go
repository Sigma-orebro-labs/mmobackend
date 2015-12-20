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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
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
	
	var buffer [1024]byte
	
	conn, err := net.ListenUDP("udp", &addr)
	
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Udp: ", ip, ":", udpPort)
	
	for {
		_, udpAddr, err := conn.ReadFromUDP(buffer[:])
		
		if err != nil {
			log.Println(err)
			continue
		}
		
		go handlers.HandleUdp(buffer, udpAddr)
	}
}