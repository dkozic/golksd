package ber

import (
	"bytes"
	"testing"
)

var bytes1 = []byte{0x62, 0x16, 0x80, 0x02, 0x0B, 0xB8, 0x82, 0x01, 0x01, 0x83, 0x02, 0xD0, 0x01, 0x88, 0x00, 0xA1, 0x04, 0x8C, 0x02, 0x01, 0x00, 0x8A, 0x01, 0x05}

func TestParseTag(t *testing.T) {

	bp := BERParser{Bytes: bytes1}
	// fmt.Printf("bp: %#v\n", bp)
	tag := bp.parseTag()

	if !bytes.Equal(tag, []byte{0x62}) {
		t.Error(`Tag != 0x62`)
	}

	if bp.Position() != 1 {
		t.Errorf(`Position() = 1, %d`, bp.Position())
	}
}

func TestParseLength(t *testing.T) {

	bp := BERParser{Bytes: bytes1}
	// fmt.Printf("bp: %#v\n", bp)

	_ = bp.parseTag()

	length, err := bp.parseLength()
	if err != nil {
		t.Error(`err != nil`)
	}
	if bp.Position() != 2 {
		t.Errorf(`Position() = 2, %d`, bp.Position())
	}
	if length != 22 {
		t.Errorf(`Length != 22, %d`, length)
	}
}

func TestParseValue(t *testing.T) {

	bp := BERParser{Bytes: bytes1}
	// fmt.Printf("bp: %#v\n", bp)

	_ = bp.parseTag()
	length, _ := bp.parseLength()

	value := bp.parseValue(length)
	if bp.Position() != 24 {
		t.Errorf(`Position() = 24, %d`, bp.Position())
	}
	if !bytes.Equal(value, []byte{0x80, 0x02, 0x0B, 0xB8, 0x82, 0x01, 0x01, 0x83, 0x02, 0xD0, 0x01, 0x88, 0x00, 0xA1, 0x04, 0x8C, 0x02, 0x01, 0x00, 0x8A, 0x01, 0x05}) {
		t.Errorf(`Value is wrong: %#v`, value)
	}
}

func TestParseTLV(t *testing.T) {

	// iteration 1
	bp := BERParser{Bytes: bytes1}
	//fmt.Printf("bp: %#v\n", bp)
	tlv, err := bp.ParseTLV()

	if err != nil {
		t.Errorf(`err != nil, %v`, err)
	}
	if !bytes.Equal(tlv.Tag, []byte{0x62}) {
		t.Errorf(`Tag != 0x62, %#v`, tlv.Tag)
	}
	if tlv.Length != 22 {
		t.Errorf(`Length != 22, %d`, tlv.Length)
	}
	if !bytes.Equal(tlv.Value, []byte{0x80, 0x02, 0x0B, 0xB8, 0x82, 0x01, 0x01, 0x83, 0x02, 0xD0, 0x01, 0x88, 0x00, 0xA1, 0x04, 0x8C, 0x02, 0x01, 0x00, 0x8A, 0x01, 0x05}) {
		t.Errorf(`Value is wrong: %#v`, tlv.Value)
	}
	if bp.HasNext() {
		t.Errorf(`HasNext() = true, %v`, bp.HasNext())
	}

	// iteration 2
	bp = BERParser{Bytes: tlv.Value}
	//fmt.Printf("bp: %#v\n", bp)
	tlv, err = bp.ParseTLV()
	if err != nil {
		t.Errorf(`err != nil, %v`, err)
	}
	if !bytes.Equal(tlv.Tag, []byte{0x80}) {
		t.Errorf(`Tag != 0x80, %#v`, tlv.Tag)
	}
	if tlv.Length != 2 {
		t.Errorf(`Length != 2, %d`, tlv.Length)
	}
	if !bytes.Equal(tlv.Value, []byte{0x0B, 0xB8}) {
		t.Errorf(`Value is wrong: %#v`, tlv.Value)
	}
	if !bp.HasNext() {
		t.Errorf(`HasNext() = false, %v`, bp.HasNext())
	}

	// iteration 3
	//fmt.Printf("bp: %#v\n", bp)
	tlv, err = bp.ParseTLV()
	if err != nil {
		t.Errorf(`err != nil, %v`, err)
	}
	if !bytes.Equal(tlv.Tag, []byte{0x82}) {
		t.Errorf(`Tag != 0x82, %#v`, tlv.Tag)
	}
	if tlv.Length != 1 {
		t.Errorf(`Length != 1, %d`, tlv.Length)
	}
	if !bytes.Equal(tlv.Value, []byte{0x01}) {
		t.Errorf(`Value is wrong: %#v`, tlv.Value)
	}
	if !bp.HasNext() {
		t.Errorf(`HasNext() = false, %v`, bp.HasNext())
	}

	// iteration 4
	//fmt.Printf("bp: %#v\n", bp)
	tlv, err = bp.ParseTLV()
	if err != nil {
		t.Errorf(`err != nil, %v`, err)
	}
	if !bytes.Equal(tlv.Tag, []byte{0x83}) {
		t.Errorf(`Tag != 0x83, %#v`, tlv.Tag)
	}
	if tlv.Length != 2 {
		t.Errorf(`Length != 2, %d`, tlv.Length)
	}
	if !bytes.Equal(tlv.Value, []byte{0xD0, 0x01}) {
		t.Errorf(`Value is wrong: %#v`, tlv.Value)
	}
	if !bp.HasNext() {
		t.Errorf(`HasNext() = false, %v`, bp.HasNext())
	}

	// iteration 5
	//fmt.Printf("bp: %#v\n", bp)
	tlv, err = bp.ParseTLV()
	if err != nil {
		t.Errorf(`err != nil, %v`, err)
	}
	if !bytes.Equal(tlv.Tag, []byte{0x88}) {
		t.Errorf(`Tag != 0x88, %#v`, tlv.Tag)
	}
	if tlv.Length != 0 {
		t.Errorf(`Length != 0, %d`, tlv.Length)
	}
	if !bytes.Equal(tlv.Value, []byte{}) {
		t.Errorf(`Value is wrong: %#v`, tlv.Value)
	}
	if !bp.HasNext() {
		t.Errorf(`HasNext() = false, %v`, bp.HasNext())
	}

	// iteration 6
	//fmt.Printf("bp: %#v\n", bp)
	tlv, err = bp.ParseTLV()
	if err != nil {
		t.Errorf(`err != nil, %v`, err)
	}
	if !bytes.Equal(tlv.Tag, []byte{0xA1}) {
		t.Errorf(`Tag != 0xA1, %#v`, tlv.Tag)
	}
	if tlv.Length != 4 {
		t.Errorf(`Length != 4, %d`, tlv.Length)
	}
	if !bytes.Equal(tlv.Value, []byte{0x8C, 0x02, 0x01, 0x00}) {
		t.Errorf(`Value is wrong: %#v`, tlv.Value)
	}
	if !bp.HasNext() {
		t.Errorf(`HasNext() = false, %v`, bp.HasNext())
	}

	// iteration 7
	//fmt.Printf("bp: %#v\n", bp)
	tlv, err = bp.ParseTLV()
	if err != nil {
		t.Errorf(`err != nil, %v`, err)
	}
	if !bytes.Equal(tlv.Tag, []byte{0x8A}) {
		t.Errorf(`Tag != 0x8A, %#v`, tlv.Tag)
	}
	if tlv.Length != 1 {
		t.Errorf(`Length != 1, %d`, tlv.Length)
	}
	if !bytes.Equal(tlv.Value, []byte{0x05}) {
		t.Errorf(`Value is wrong: %#v`, tlv.Value)
	}
	if bp.HasNext() {
		t.Errorf(`HasNext() = true, %v`, bp.HasNext())
	}
}
