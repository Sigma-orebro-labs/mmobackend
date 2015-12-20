#include "udp_client.h"

UdpClient::UdpClient(const char* ip_address, USHORT port)
{
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

void UdpClient::send()
{
	auto message = "Hello from client!";
	//send the message
	if (sendto(socket_, message, strlen(message), 0, (struct sockaddr*)&sockaddr_in_, sizeof(sockaddr_in_)) == SOCKET_ERROR)
	{
		printf("sendto() failed with error code : %d", WSAGetLastError());
	}
}
