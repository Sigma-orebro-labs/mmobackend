#pragma once

#define WIN32_LEAN_AND_MEAN

#include <windows.h>
#include <winsock2.h>
#include <ws2tcpip.h>

#include "../exceptions/mmo_exception.h"
#include "message.h"

class UdpClient
{
public:
	UdpClient(const char* ip_address, USHORT port);
	~UdpClient();

	void send(Message& msg);

private:
	SOCKET socket_;
	struct sockaddr_in sockaddr_in_;
};
