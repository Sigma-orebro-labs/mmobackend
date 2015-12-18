package util

import (
	"testing"
)

func TestBytesToUint16(t *testing.T) {
	v := BytesToUint16(255, 255)
	if v != 65535 {
		t.Fail()
	}
	
	v = BytesToUint16(0, 0)
	if v != 0 {
		t.Fail()
	}
	
	v = BytesToUint16(128, 0)
	if v != 128 {
		t.Fail()
	}
	
	v = BytesToUint16(89, 67)
	if v != 17241 {
		t.Fail()
	}
}

func TestUint16ToBytes(t *testing.T) {
	b1, b2 := Uint16ToBytes(65535)
	if b1 != 255 || b2 != 255 {
		t.Fail()
	}
	
	b1, b2 = Uint16ToBytes(0)
	if b1 != 0 || b2 != 0 {
		t.Fail()
	}
	
	b1, b2 = Uint16ToBytes(255)
	if b1 != 255 || b2 != 0 {
		t.Fail()
	}
	
	b1, b2 = Uint16ToBytes(56567)
	if b1 != 247 || b2 != 220 {
		t.Fail()
	}        
}
