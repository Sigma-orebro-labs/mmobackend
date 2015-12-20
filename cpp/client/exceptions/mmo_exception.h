#pragma once

#include <exception>
#include <string>

class MmoException : public std::exception
{
public:
	MmoException(const std::string& message) : message_(message) {};

	virtual const char* what() const throw()
	{
		return message_.c_str();
	}

protected:
	std::string message_;
};
