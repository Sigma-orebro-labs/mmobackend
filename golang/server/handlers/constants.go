package handlers

const (
	maxMessageBodyLength byte = 128
	messageHeaderMarker byte = 0xF0 	// 1111 0000
	messageFooterMarker byte = 0xCC 	// 1100 1100
	messageHeaderLength byte = 4		// The last 2 bytes are used for the content length
	messageFooterLength byte = 1
	
	// Commands
	getCurrentUserPositionCommandCode byte = 1
	getCurrentUserPositionResponseCode byte = 2
	getEnemyPositionsCommandCode byte = 3
	getEnemyPositionsResponseCode byte = 4
)