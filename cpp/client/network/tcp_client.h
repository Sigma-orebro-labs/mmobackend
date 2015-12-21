#pragma once

#define WIN32_LEAN_AND_MEAN

#include <windows.h>
#include <winsock2.h>
#include <ws2tcpip.h>

#include "../exceptions/mmo_exception.h"

class TcpClient
{
public:
	TcpClient(const char* ip_address, const char* port);
	~TcpClient();

	bool create_connection();

private:
	SOCKET socket_;
	struct addrinfo* result_;
	struct addrinfo* ptr_;
	struct addrinfo hints_;
};