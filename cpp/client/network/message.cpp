#include "message.h"

Message::Message(Command c)
	:
	index_(0)
{
	message_[0] = HEADER;
	message_[1] = static_cast<char>(c);

	// We don't know the content length yet
	message_[2] = 0;
	message_[3] = 0;
}

void Message::append_to_body(char v)
{
	message_[index_++] = v;
}

void Message::append_to_body(uint16_t v)
{
	char b1, b2;
	uint16_to_bytes(v, b1, b2);
	message_[index_++] = b1;
	message_[index_++] = b2;
}

const char* Message::bytes()
{
	char b1, b2;
	uint16_to_bytes(index_, b1, b2);

	message_[2] = b1;
	message_[3] = b2;

	return message_;
}

int Message::length() const
{
	return MESSAGE_HEADER_LENGTH + index_ + MESSAGE_FOOTER_LENGTH;
}
