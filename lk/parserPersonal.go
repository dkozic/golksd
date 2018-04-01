package lk

import "github.com/dkozic/golksd/ber"

type parserPersonal struct{}

func newParserPersonal() parserPersonal {
	p := parserPersonal{}
	return p
}

func (p parserPersonal) parseFile(bytes []byte) (map[TagPersonal]ber.TLV, error) {
	result := make(map[TagPersonal]ber.TLV)
	bp := ber.BERParser{Bytes: bytes, Skip00: true, Skip06: true}
	for i := 0; i < 10; i++ {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGP_0D.code:
			result[TAGP_0D] = tlv
		case TAGP_16.code:
			result[TAGP_16] = tlv
		case TAGP_17.code:
			result[TAGP_17] = tlv
		case TAGP_18.code:
			result[TAGP_18] = tlv
		case TAGP_19.code:
			result[TAGP_19] = tlv
		case TAGP_1A.code:
			result[TAGP_1A] = tlv
		case TAGP_1B.code:
			result[TAGP_1B] = tlv
		case TAGP_1D.code:
			result[TAGP_1D] = tlv
		case TAGP_1E.code:
			result[TAGP_1E] = tlv
		case TAGP_1F06.code:
			result[TAGP_1F06] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}
