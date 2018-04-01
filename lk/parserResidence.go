package lk

import "github.com/dkozic/golksd/ber"

type parserResidence struct{}

func newParserResidence() parserResidence {
	p := parserResidence{}
	return p
}

func (p parserResidence) parseFile(bytes []byte) (map[TagResidence]ber.TLV, error) {
	result := make(map[TagResidence]ber.TLV)
	bp := ber.BERParser{Bytes: bytes, Skip00: true, Skip06: true}
	for i := 0; i < 7; i++ {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGR_0D.code:
			result[TAGR_0D] = tlv
		case TAGR_20.code:
			result[TAGR_20] = tlv
		case TAGR_21.code:
			result[TAGR_21] = tlv
		case TAGR_22.code:
			result[TAGR_22] = tlv
		case TAGR_23.code:
			result[TAGR_23] = tlv
		case TAGR_24.code:
			result[TAGR_24] = tlv
		case TAGR_2A.code:
			result[TAGR_2A] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}
