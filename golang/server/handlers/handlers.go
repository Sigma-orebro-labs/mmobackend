package handlers

import (
	"net"
	"log"
	"io"
	"github.com/Sigma-orebro-labs/mmobackend/golang/server/util"
)

const (
	maxMessageBodyLength byte = 128;
	messageHeaderMarker byte = 0xF0 	// 1111 0000
	messageFooterMarker byte = 0xCC 	// 1100 1100
	messageHeaderLength byte = 4		// The last 2 bytes are used for the content length
	messageFooterLength byte = 1
	
	// Commands
	getCurrentUserPositionCommandCode byte = 1
	getCurrentUserPositionResponseCode byte = 2
	getEnemyPositionsCommandCode byte = 3;
	getEnemyPositionsResponseCode byte = 4;
)

func Handle(conn net.Conn) {
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
		case getEnemyPositionsCommandCode:
			getEnemyPositionsCommand(response, conn, responseBaseSize)
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

func getEnemyPositionsCommand(response []byte, conn net.Conn, responseBaseSize byte) {
	response[1] = getEnemyPositionsResponseCode;
	
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
				