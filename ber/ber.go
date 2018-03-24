package ber

import "fmt"

type TLV struct {
	Tag    []byte
	Length int
	Value  []byte
}

type BERParser struct {
	next   int
	Bytes  []byte
	Skip00 bool
	Skip06 bool
}

func (p *BERParser) Position() int {
	return p.next
}

func (p *BERParser) HasNext() bool {
	return p.next < len(p.Bytes)
}

func (p *BERParser) ParseTLV() (tlv TLV, err error) {
	tlv.Tag = p.parseTag()
	tlv.Length, err = p.parseLength()
	if err != nil {
		return tlv, err
	}
	tlv.Value = p.parseValue(tlv.Length)

	return tlv, nil

}

func (p *BERParser) parseTag() []byte {
	temp := p.next

	res := []byte{}
	for p.Bytes[temp] == 0x00 {
		temp++
	}
	byte1 := p.Bytes[temp]
	temp++
	if (byte1 & 0x1F) != 0x1F {
		res = append(res, byte1&0xFF)

	} else {
		res = append(res, byte1&0xFF)

		for {
			b := p.Bytes[temp]
			temp++
			res = append(res, b&0xFF)

			if (b & 0x80) != 0x80 {
				break
			}
		}

	}

	p.next = temp
	return res

}

func (p *BERParser) parseLength() (int, error) {
	temp := p.next
	if p.next >= len(p.Bytes) {
		return 0, fmt.Errorf("BerParserError: You are at the end of byte array at position: %d, %d, %#v", p.next, temp, p.Bytes)
	}

	if p.Skip06 && p.Bytes[temp] == 0x06 {
		temp++
	}

	length := 0
	var byte1 byte = p.Bytes[temp]
	temp++
	if (byte1 & 0x80) != 0x80 {
		length = int(byte1 & 0xFF)
	} else {
		var l int = int(byte1 & 0x7F)
		for i := 0; i < l; i++ {
			length = length + (int(p.Bytes[temp+i]&0xFF) << (8 * uint(l-1-i)))
		}
		temp = temp + l
	}
	p.next = temp
	return length, nil
}

func (p *BERParser) parseValue(length int) []byte {

	res := []byte{}
	for p.Skip00 && p.Bytes[p.next] == 0x00 {
		p.next = p.next + 1
	}

	for i := 0; i < length; i++ {
		res = append(res, p.Bytes[p.next])
		p.next = p.next + 1
	}
	return res

}
