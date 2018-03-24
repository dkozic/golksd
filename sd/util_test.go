package sd

import (
	"testing"
)

func TestBytesToHex1(t *testing.T) {

	s := bytesToHex([]byte{0x01})

	if s != "01" {
		t.Errorf(`res != 01, %s`, s)
	}

}

func TestBytesToHex2(t *testing.T) {

	s := bytesToHex([]byte{0xA1, 0x43})

	if s != "A143" {
		t.Errorf(`res != A143, %s`, s)
	}

}
