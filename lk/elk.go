package lk

import (
	"fmt"
	"log"

	"github.com/dkozic/golksd/ber"
	"github.com/ebfe/scard"
)

var EF_DOCUMENT = []byte{0x0F, 0x02}
var EF_PERSONAL = []byte{0x0F, 0x03}
var EF_RESIDENCE = []byte{0x0F, 0x04}
var EF_PORTRAIT = []byte{0x0F, 0x06}

func ReadElk(card *scard.Card) (ElkData, error) {

	data := ElkData{}

	if err := initCard(card); err != nil {
		return data, fmt.Errorf("error doing initCard() in ReadElk(): %v", err)
	}

	fileDocument, err := readFile(card, EF_DOCUMENT)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(DOCUMENT) in ReadElk(): %v", err)
	}
	fieldsDocument, err := newParserDocument().parseFile(fileDocument)
	if err != nil {
		return data, fmt.Errorf("error doing parserDocument.parseFile() in ReadElk(): %v", err)
	}
	//log.Printf("fieldsDocument: %#v\n", fieldsDocument)

	filePersonal, err := readFile(card, EF_PERSONAL)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(PERSONAL) in ReadElk(): %v", err)
	}
	fieldsPersonal, err := newParserPersonal().parseFile(filePersonal)
	if err != nil {
		return data, fmt.Errorf("error doing parserPersonal.parseFile() in ReadElk(): %v", err)
	}
	//log.Printf("fieldsPersonal: %#v\n", fieldsPersonal)

	fileResidence, err := readFile(card, EF_RESIDENCE)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(RESIDENCE) in ReadElk(): %v", err)
	}
	fieldsResidence, err := newParserResidence().parseFile(fileResidence)
	if err != nil {
		return data, fmt.Errorf("error doing parserResidence.parseFile() in ReadElk(): %v", err)
	}

	filePortrait, err := readFile(card, EF_PORTRAIT)
	if err != nil {
		return data, fmt.Errorf("error doing readFile(PORTRAIT) in ReadElk(): %v", err)
	}
	portrait := filePortrait

	f := newDataFactory(fieldsDocument, fieldsPersonal, fieldsResidence, portrait)
	data = f.create()
	// log.Printf("data: %#v\n", data)

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

func selectFile(card *scard.Card, fileId []byte) ([]byte, error) {

	//command := []byte{0x00, 0xA4, 0x02, 0x04}
	command := []byte{0x00, 0xA4, 0x08, 0x00}
	command = append(command, byte(len(fileId)))
	command = append(command, fileId...)

	printCommand("SELECT FILE", command)

	response, err := card.Transmit(command)
	if err != nil {
		return nil, fmt.Errorf("error in selectFile: %v", err)
	}

	printResponse(response)

	return response[:len(response)-2], nil
}

func stripHeader(file []byte, offset int) []byte {
	return file[offset:]
}

func readFile(card *scard.Card, fileId []byte) ([]byte, error) {

	// select
	fcp, err := selectFile(card, fileId)
	if err != nil {
		return nil, fmt.Errorf("error doing selectFile() in readFile(): %v", err)
	}

	parserFCP := newParserFCP()

	log.Printf("fcp: %#v", fcp)

	fcpTags, err := parserFCP.parseFCP(fcp)

	tlv, ok := fcpTags[TAGF_81]
	if !ok {
		return nil, fmt.Errorf("there is no tag 81 in FCP")
	}

	fileLength := parseFileLength(tlv)

	res := []byte{}

	err = card.BeginTransaction()
	if err != nil {
		return nil, fmt.Errorf("error calling card.BeginTransaction()")
	}

	defer card.EndTransaction(scard.LeaveCard)

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

func parseFileLength(tlv ber.TLV) int {
	log.Printf("tag: %#v\n", tlv)
	v2 := tlv.Value
	l2 := tlv.Length
	length := 0
	for i := 0; i < l2; i++ {
		length = length + (int(v2[i]&0xFF) << (8 * uint(l2-1-i)))
	}
	log.Printf("length: %d\n", length)
	return length
}

func printCommand(command string, apdu []byte) {
	log.Printf("%s APDU >>>: %#v\n", command, apdu)
}

func printResponse(apdu []byte) {
	log.Printf("RESPONSE APDU <<<: %#v\n", apdu)
}
