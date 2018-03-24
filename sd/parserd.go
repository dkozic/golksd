package sd

import "github.com/dkozic/golksd/ber"

type parserD struct{}

func newParserD() parserD {
	p := parserD{}
	return p
}

func (p parserD) parse78(bytes []byte) (map[TagD]ber.TLV, error) {
	result := make(map[TagD]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGD_4F.code:
			result[TAGD_4F] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserD) parse72(bytes []byte) (map[TagD]ber.TLV, error) {
	result := make(map[TagD]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGD_80.code:
			result[TAGD_80] = tlv
		case TAGD_C9.code:
			result[TAGD_C9] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserD) parseFile(bytes []byte) (map[TagD]ber.TLV, error) {
	result := make(map[TagD]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for i := 0; i < 2; i++ {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGD_78.code:
			a, err := p.parse78(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGD_72.code:
			a, err := p.parse72(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}
