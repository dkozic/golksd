package lk

import "github.com/dkozic/golksd/ber"

type parserDocument struct{}

func newParserDocument() parserDocument {
	p := parserDocument{}
	return p
}

func (p parserDocument) parseFile(bytes []byte) (map[TagDocument]ber.TLV, error) {
	result := make(map[TagDocument]ber.TLV)
	bp := ber.BERParser{Bytes: bytes, Skip00: true, Skip06: true}
	for i := 0; i < 10; i++ {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGD_0D.code:
			result[TAGD_0D] = tlv
		case TAGD_0A.code:
			result[TAGD_0A] = tlv
		case TAGD_0B.code:
			result[TAGD_0B] = tlv
		case TAGD_0C.code:
			result[TAGD_0C] = tlv
		case TAGD_0D.code:
			result[TAGD_0D] = tlv
		case TAGD_0E.code:
			result[TAGD_0E] = tlv
		case TAGD_0F.code:
			result[TAGD_0F] = tlv
		case TAGD_09.code:
			result[TAGD_09] = tlv
		case TAGD_10.code:
			result[TAGD_10] = tlv
		case TAGD_11.code:
			result[TAGD_11] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}
