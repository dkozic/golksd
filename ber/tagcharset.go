package ber

import (
	"fmt"

	"golang.org/x/text/encoding/charmap"
)

type TagCharset string

const HEX TagCharset = "HEX"
const ISO8859P1 TagCharset = "ISO8859-1"
const ISO8859P2 TagCharset = "ISO8859-2"
const ISO8859P5 TagCharset = "ISO8859-5"
const ISO8859P7 TagCharset = "ISO8859-7"
const UTF8 TagCharset = "UTF-8"

func (cs TagCharset) BytesToCharsetString(bytes []byte) string {
	res := ""
	switch cs {
	case HEX:
		res = fmt.Sprintf("%X", bytes)
	case ISO8859P1:
		bs, err := charmap.ISO8859_1.NewDecoder().Bytes(bytes)
		if err != nil {
			fmt.Printf("error converting to ISO8859P1: %#v", bytes)
		} else {
			res = string(bs)
		}
	case ISO8859P2:
		bs, err := charmap.ISO8859_2.NewDecoder().Bytes(bytes)
		if err != nil {
			fmt.Printf("error converting to ISO8859P2: %#v", bytes)
		} else {
			res = string(bs)
		}
	case ISO8859P5:
		bs, err := charmap.ISO8859_5.NewDecoder().Bytes(bytes)
		if err != nil {
			fmt.Printf("error converting to ISO8859P5: %#v", bytes)
		} else {
			res = string(bs)
		}
	case ISO8859P7:
		bs, err := charmap.ISO8859_2.NewDecoder().Bytes(bytes)
		if err != nil {
			fmt.Printf("error converting to ISO8859P7: %#v", bytes)
		} else {
			res = string(bs)
		}
	case UTF8:
		res = string(bytes)
	default:
		res = string(bytes)
	}

	return res

}
