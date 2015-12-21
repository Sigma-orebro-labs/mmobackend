#pragma once

#include <stdint.h>
#include <vector>

#include "../util/type_converter.h"

const std::size_t MESSAGE_HEADER_LENGTH = 4;	// The last 2 bytes are used for the content length
const std::size_t MESSAGE_FOOTER_LENGTH = 1;
const std::size_t MAX_MESSAGE_BODY_LENGTH = 128;

const unsigned char HEADER = 0xF0;
const unsigned char FOOTER = 0xCC;

// Commands
const char GET_CURRENT_USER_POSITION_COMMAND = 1;
const char GET_CURRENT_USER_POSITION_RESPONSE = 2;
const char UPDATE_POSITION_COMMAND = 3;
const char UPDATE_ENEMY_POSITIONS_COMMAND = 4;

enum class Command : char
{
	GET_CURRENT_USER_POSITION_COMMAND = 1,
	GET_CURRENT_USER_POSITION_RESPONSE = 2,
	UPDATE_POSITION_COMMAND = 3,
	UPDATE_ENEMY_POSITIONS_COMMAND = 4,
};

class Message
{
public:
	Message(Command c);
	void append_to_body(char v);
	void append_to_body(uint16_t v);

	const char* bytes();
	int length() const;
private:
	char message_[MESSAGE_HEADER_LENGTH + MAX_MESSAGE_BODY_LENGTH + MESSAGE_FOOTER_LENGTH];
	uint16_t index_;
};
