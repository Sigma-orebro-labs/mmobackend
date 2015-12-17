#include "client.h"

Client::Client()
{
	socket_ = INVALID_SOCKET;
	result_ = nullptr;
	ptr_ = nullptr;
}

bool Client::init()
{
	int result;

	// Initialize Winsock
	result = WSAStartup(MAKEWORD(2, 2), &wsa_data_);
	if (result != 0)
	{
		printf("WSAStartup failed with error: %d\n", result);
		return 1;
	}

	ZeroMemory(&hints_, sizeof(hints_));
	hints_.ai_family = AF_UNSPEC;
	hints_.ai_socktype = SOCK_STREAM;
	hints_.ai_protocol = IPPROTO_TCP;

	// Resolve the server address and port
	result = getaddrinfo("localhost", "11000", &hints_, &result_);
	if (result != 0)
	{
		printf("getaddrinfo failed with error: %d\n", result);
		WSACleanup();
		return false;
	}

	return true;
}

bool Client::create_connection()
{
	int result;

	// Attempt to connect to an address until one succeeds
	for (ptr_ = result_; ptr_ != nullptr; ptr_ = ptr_->ai_next) {

		// Create a socket for connecting to the server
		socket_ = socket(ptr_->ai_family, ptr_->ai_socktype, ptr_->ai_protocol);
		if (socket_ == INVALID_SOCKET) 
		{
			printf("socket failed with error: %ld\n", WSAGetLastError());
			WSACleanup();
			return false;
		}

		// Connect to the server.
		result = connect(socket_, ptr_->ai_addr, (int)ptr_->ai_addrlen);
		if (result == SOCKET_ERROR)
		{
			closesocket(socket_);
			socket_ = INVALID_SOCKET;
			continue;
		}
		break;
	}

	freeaddrinfo(result_);

	if (socket_ == INVALID_SOCKET) 
	{
		printf("Unable to connect to the server!\n");
		WSACleanup();
		return false;
	}

	return true;
}

void Client::close_connection()
{
	closesocket(socket_);
	WSACleanup();
}

void Client::get_enemy_positions(int& x, int& y)
{
	const int REQUEST_BUFFER_LENGTH = MESSAGE_HEADER_LENGTH + MESSAGE_FOOTER_LENGTH + 1;
	char request_buffer[REQUEST_BUFFER_LENGTH];

	uint16_t contentLength = 1;
	request_buffer[0] = MESSAGE_HEADER_MARKER;
	request_buffer[1] = GET_ENEMY_POSITION_COMMAND;
	request_buffer[2] = (char)(contentLength & 0xFF);
	request_buffer[3] = (char)(contentLength & 0xFF);
	request_buffer[4] = 123;
	request_buffer[5] = MESSAGE_FOOTER_MARKER;

	const int RESPONSE_BUFFER_LENGTH = MESSAGE_HEADER_LENGTH + MESSAGE_FOOTER_LENGTH + 3;
	char response_buffer[MESSAGE_HEADER_LENGTH + MESSAGE_FOOTER_LENGTH + 3];

	int result;

	// Send an initial buffer
	result = send(socket_, request_buffer, REQUEST_BUFFER_LENGTH, 0);
	if (result == SOCKET_ERROR) 
	{
		printf("send failed with error: %d\n", WSAGetLastError());
		closesocket(socket_);
		WSACleanup();
		return;
	}

	printf("Bytes Sent: %ld\n", result);

	result = recv(socket_, response_buffer, RESPONSE_BUFFER_LENGTH, 0);
	if (result > 0)
	{
		printf("Bytes received: %d\n", result);
		if (response_buffer[1] == GET_ENEMY_POSITION_RESPONSE)
		{
			printf("Recived %i \n", response_buffer[1]);
		}
		else
		{
			printf("Unexpected response code \n");
			return;
		}

		x = response_buffer[4]; 	// x
		y = response_buffer[5]; 	// y

		printf("X: %i \n", response_buffer[4]);
		printf("Y: %i \n", response_buffer[5]);
	}
	else if (result == 0)
		printf("Connection closed\n");
	else
		printf("recv failed with error: %d\n", WSAGetLastError());
}
