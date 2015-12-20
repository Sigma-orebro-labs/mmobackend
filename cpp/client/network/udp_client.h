#pragma once

#include "client.h"

class UdpClient : public Client
{
public:
	UdpClient(const char* ip_address, USHORT port);
	~UdpClient();

	void send();

private:
	struct sockaddr_in sockaddr_in_;
};
