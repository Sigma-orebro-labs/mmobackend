package handlers

import (
	"net"
	"log"
	"io"
	"github.com/Sigma-orebro-labs/mmobackend/golang/server/util"
	//"github.com/Sigma-orebro-labs/mmobackend/golang/server/game"
)

func HandleUdp(data [1024]byte, addr *net.UDPAddr) {
	if data[0] != messageHeaderMarker {
		log.Println("Invalid message header")
	}
	
	commandCode := data[1];
	bodyLength := util.BytesToUint16(data[2], data[3])
	
	footer := data[uint16(messageHeaderLength) + bodyLength]	
	if footer != messageFooterMarker {
		log.Fatal("Invalid message footer")
	}
	
	switch (commandCode) {
	case updatePlayerPositionCommandCode:
		x := util.BytesToUint16(data[4], data[5])
		y := util.BytesToUint16(data[6], data[7])
		log.Println("X: ", x)
		log.Println("Y: ", y)
	    break;
	default:
	    break;
	}
}

func HandleTcp(conn net.Conn) {
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
		bodyLength := util.BytesToUint16(header[2], header[3])
		log.Println("bodyLength: ", bodyLength)
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
		case getCurrentUserPositionCommandCode:
			getCurrentUserPositionCommand(response, conn, responseBaseSize)							
		    break;				
		default:
		    break;
		}
	}
}

func getCurrentUserPositionCommand(response []byte, conn net.Conn, responseBaseSize byte) {
	response[1] = getCurrentUserPositionResponseCode;
	
	// 3 byte response body (x, y, z coordinates)
	b1, b2 := util.Uint16ToBytes(3)
	response[2] = b1; 	
	response[3] = b2;
	 	
	response[4] = 101; 	// x
	response[5] = 102; 	// y
	response[6] = 103; 	// z
	response[7] = messageFooterMarker;
	
	_, err := conn.Write(response[0:responseBaseSize + 3])	
	if err != nil {
		log.Println(err)
	}
}
				