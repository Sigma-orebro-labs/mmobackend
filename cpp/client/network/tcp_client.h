#pragma once

#include "client.h"

class TcpClient : public Client
{
public:
	TcpClient(const char* ip_address, const char* port);
	~TcpClient();

	bool create_connection();
	void get_enemy_positions(int& x, int& y);

private:
	struct addrinfo* result_;
	struct addrinfo* ptr_;
	struct addrinfo hints_;
};