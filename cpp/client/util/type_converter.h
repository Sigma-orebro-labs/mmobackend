#pragma once

#include <stdint.h>

inline uint16_t bytes_to_uint16(char b1, char b2)
{
	return (uint16_t)((b2 << 8) | b1);
}

inline void uint16_to_bytes(uint16_t number, char& b1, char& b2)
{
	b2 = (char)(number >> 8);
	b1 = (char)number;
}
