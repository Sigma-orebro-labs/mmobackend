#include "udp_client.h"

UdpClient::UdpClient(const char* ip_address, USHORT port)
{
	socket_ = INVALID_SOCKET;

	if ((socket_ = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP)) == SOCKET_ERROR)
	{
		printf("socket() failed with error code : %d", WSAGetLastError());
		throw MmoException("getaddrinfo failed");
	}

	memset((char *)&sockaddr_in_, 0, sizeof(sockaddr_in_));
	sockaddr_in_.sin_family = AF_INET;
	sockaddr_in_.sin_port = htons(port);
	inet_pton(AF_INET, ip_address, &sockaddr_in_.sin_addr.S_un.S_addr);
}

UdpClient::~UdpClient()
{
	closesocket(socket_);
}

void UdpClient::send(Message& msg)
{
	/*auto message = "Hello from client!";
	if (sendto(socket_, message, strlen(message), 0, (struct sockaddr*)&sockaddr_in_, sizeof(sockaddr_in_)) == SOCKET_ERROR)
	{
		printf("sendto() failed with error code : %d", WSAGetLastError());
	}*/

	if (sendto(socket_, msg.bytes(), msg.length(), 0, (struct sockaddr*)&sockaddr_in_, sizeof(sockaddr_in_)) == SOCKET_ERROR)
	{
		printf("UDP send failed with error code : %d", WSAGetLastError());
	}
}
