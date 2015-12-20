#include "win_sock.h"

bool WinSock::initialize()
{
	WSADATA wsa_data;
	if (WSAStartup(MAKEWORD(2, 2), &wsa_data) != 0)
	{
		printf("WSAStartup failed with error: %d\n", WSAGetLastError());
		return false;
	}

	return true;
}

void WinSock::cleanup()
{
	WSACleanup();
}
