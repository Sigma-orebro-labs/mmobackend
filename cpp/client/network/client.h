#pragma once

#define WIN32_LEAN_AND_MEAN

#include <windows.h>
#include <winsock2.h>
#include <ws2tcpip.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>

// Link with Ws2_32.lib, Mswsock.lib, and Advapi32.lib
#pragma comment (lib, "Ws2_32.lib")
#pragma comment (lib, "Mswsock.lib")
#pragma comment (lib, "AdvApi32.lib")

const char MAX_MESSAGE_BODY_LENGTH = 128;
const char MESSAGE_HEADER_MARKER = 0xF0; 	// 1111 0000
const char MESSAGE_FOOTER_MARKER = 0xCC; 	// 1100 1100
const char MESSAGE_HEADER_LENGTH = 4;		// The last 2 bytes are used for the content length
const char MESSAGE_FOOTER_LENGTH = 1;

// Commands
const char GET_CURRENT_USER_POSITION_COMMAND = 1;
const char GET_CURRENT_USER_POSITION_RESPONSE = 2;
const char GET_ENEMY_POSITION_COMMAND = 3;
const char GET_ENEMY_POSITION_RESPONSE = 4;

class Client
{
public:
	Client();

	bool init();
	bool create_connection();
	void close_connection();
	void get_enemy_positions(int& x, int& y);

private:
	WSADATA wsa_data_;
	SOCKET socket_;
	struct addrinfo* result_;
	struct addrinfo* ptr_;
	struct addrinfo hints_;
};
