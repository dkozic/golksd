package sd

import (
	"bytes"
	"fmt"
	"log"

	"github.com/dkozic/golksd/ber"
	"github.com/ebfe/scard"
)

var APPLICATION = []byte{0xA0, 0x00, 0x00, 0x00, 0x77, 0x01, 0x08, 0x00, 0x07, 0x00, 0x00, 0xFE, 0x00, 0x00, 0xAD, 0xF2}

var EF_REGISTRATON_A = []byte{0xD0, 0x01}
var EF_REGISTRATON_B = []byte{0xD0, 0x11}
var EF_REGISTRATON_C = []byte{0xD0, 0x21}
var EF_REGISTRATON_D = []byte{0xD0, 0x31}

func ReadEsd(card *scard.Card) (EsdData, error) {

	data := EsdData{}
	if err := initCard(card); err != nil {
		return data, fmt.Errorf("error doing initCard() in ReadEsd(): %v", err)
	}

	if err := selectApplication(card); err != nil {
		return data, fmt.Errorf("error doing initCard() in ReadEsd(): %v", err)
	}

	fileA, err := readFile(card, EF_REGISTRATON_A)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(A) in ReadEsd(): %v", err)
	}
	fieldsA, err := newParserA().parseFile(fileA)
	if err != nil {
		return data, fmt.Errorf("error doing parserA.parseFile() in ReadEsd(): %v", err)
	}
	// log.Printf("fieldsA: %#v", fieldsA)

	fileB, err := readFile(card, EF_REGISTRATON_B)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(B) in ReadEsd(): %v", err)
	}
	fieldsB, err := newParserB().parseFile(fileB)
	if err != nil {
		return data, fmt.Errorf("error doing parserB.parseFile() in ReadEsd(): %v", err)
	}
	// log.Printf("fieldsB: %#v", fieldsB)

	fileC, err := readFile(card, EF_REGISTRATON_C)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(C) in ReadEsd(): %v", err)
	}
	fieldsC, err := newParserC().parseFile(fileC)
	if err != nil {
		return data, fmt.Errorf("error doing parserC.parseFile() in ReadEsd(): %v", err)
	}

	fileD, err := readFile(card, EF_REGISTRATON_D)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(D) in ReadEsd(): %v", err)
	}
	fieldsD, err := newParserD().parseFile(fileD)
	if err != nil {
		return data, fmt.Errorf("error doing parserD.parseFile() in ReadEsd(): %v", err)
	}

	f := newDataFactory(fieldsA, fieldsB, fieldsC, fieldsD)
	data = f.create()
	// log.Printf("data: %#v", data)

	return data, nil

}

func initCard(card *scard.Card) error {
	// init
	log.Printf("Card: %v\n", card)

	status, err := card.Status()
	if err != nil {
		return fmt.Errorf("error in init: %v", err)
	}

	// atr
	atr := status.Atr
	log.Printf("ATR: %#v\n", atr)
	return nil
}

func selectApplication(card *scard.Card) error {

	command := []byte{0x00, 0xA4, 0x04, 0x0C}
	command = append(command, byte(len(APPLICATION)))
	command = append(command, APPLICATION...)

	printCommand("SELECT APPLICATION", command)

	response, err := card.Transmit(command)
	if err != nil {
		return fmt.Errorf("error in selectApplication: %v", err)
	}

	printResponse(response)
	return nil

}

func selectFile(card *scard.Card, fileId []byte) (int, error) {

	command := []byte{0x00, 0xA4, 0x02, 0x04}
	command = append(command, byte(len(fileId)))
	command = append(command, fileId...)

	printCommand("SELECT FILE", command)

	response, err := card.Transmit(command)
	if err != nil {
		return 0, fmt.Errorf("error in selectFile: %v", err)
	}

	printResponse(response)

	length, err := parseFileLength(response[:len(response)-2])
	if err != nil {
		return 0, fmt.Errorf("error in selectFile: %v", err)
	}
	return length, nil
}

func readFile(card *scard.Card, fileId []byte) ([]byte, error) {

	// select
	fileLength, err := selectFile(card, fileId)
	if err != nil {
		return nil, fmt.Errorf("error doing selectFile() in readFile(): %v", err)
	}

	res := []byte{}

	// read binary
	offset := 0
	for {
		var p1 int = (offset >> 8)
		var p2 int = (offset & 0x00FF)
		command := []byte{0x00, 0xB0, byte(p1), byte(p2), 0}
		printCommand("READ BYNARY", command)

		response, err := card.Transmit(command)
		if err != nil {
			return nil, fmt.Errorf("error in readFile: %v", err)
		}

		printResponse(response)

		response = response[:len(response)-2]
		res = append(res, response...)
		offset = offset + len(response)
		if offset >= fileLength {
			break
		}
	}

	return res, nil
}

func parseFileLength(fcp []byte) (int, error) {
	bp1 := ber.BERParser{Bytes: fcp}
	tlv1, err := bp1.ParseTLV()
	if err != nil {
		return 0, err
	}

	bytes2 := tlv1.Value
	bp2 := ber.BERParser{Bytes: bytes2}
	tlv2, err := bp2.ParseTLV()
	if err != nil {
		return 0, err
	}

	if !bytes.Equal(tlv2.Tag, []byte{0x80}) {
		return 0, fmt.Errorf("first tag is not 0x80: %#v", tlv2.Tag)
	}
	v2 := tlv2.Value
	l2 := tlv2.Length
	length := 0
	for i := 0; i < l2; i++ {
		length = length + (int(v2[i]&0xFF) << (8 * uint(l2-1-i)))
	}
	return length, nil

}

func printCommand(command string, apdu []byte) {
	log.Printf("%s APDU >>>: %#v\n", command, apdu)
}

func printResponse(apdu []byte) {
	log.Printf("RESPONSE APDU <<<: %#v\n", apdu)
}
