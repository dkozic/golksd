package sd

import "github.com/dkozic/golksd/ber"

type parserC struct{}

func newParserC() parserC {
	p := parserC{}
	return p
}

func (p parserC) parse78(bytes []byte) (map[TagC]ber.TLV, error) {
	result := make(map[TagC]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGC_4F.code:
			result[TAGC_4F] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserC) parse72(bytes []byte) (map[TagC]ber.TLV, error) {
	result := make(map[TagC]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGC_80.code:
			result[TAGC_80] = tlv
		case TAGC_C1.code:
			result[TAGC_C1] = tlv
		case TAGC_C2.code:
			result[TAGC_C2] = tlv
		case TAGC_C3.code:
			result[TAGC_C3] = tlv
		case TAGC_C4.code:
			result[TAGC_C4] = tlv
		case TAGC_C5.code:
			result[TAGC_C5] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserC) parseFile(bytes []byte) (map[TagC]ber.TLV, error) {
	result := make(map[TagC]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for i := 0; i < 2; i++ {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGC_78.code:
			a, err := p.parse78(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGC_72.code:
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
