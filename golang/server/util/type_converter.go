package util

func BytesToUint16(b1, b2 byte) uint16 {
	return (uint16(b2) << 8) | uint16(b1);
}

func Uint16ToBytes(number uint16) (b1, b2 byte) {
	return uint8(number & 0xff), uint8(number >> 8)
}