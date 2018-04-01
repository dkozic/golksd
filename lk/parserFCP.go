package lk

import "github.com/dkozic/golksd/ber"

type parserFCP struct{}

func newParserFCP() parserFCP {
	p := parserFCP{}
	return p
}

func (p parserFCP) parse62(bytes []byte) (map[TagFCP]ber.TLV, error) {
	result := make(map[TagFCP]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGF_80.code:
			result[TAGF_80] = tlv
		case TAGF_81.code:
			result[TAGF_81] = tlv
		case TAGF_82.code:
			result[TAGF_82] = tlv
		case TAGF_83.code:
			result[TAGF_83] = tlv
		case TAGF_84.code:
			result[TAGF_84] = tlv
		case TAGF_85.code:
			result[TAGF_85] = tlv
		case TAGF_86.code:
			result[TAGF_86] = tlv
		case TAGF_87.code:
			result[TAGF_87] = tlv
		case TAGF_88.code:
			result[TAGF_88] = tlv
		case TAGF_8A.code:
			result[TAGF_8A] = tlv
		case TAGF_8B.code:
			result[TAGF_8B] = tlv
		case TAGF_8C.code:
			result[TAGF_8C] = tlv
		case TAGF_8D.code:
			result[TAGF_8D] = tlv
		case TAGF_8E.code:
			result[TAGF_8E] = tlv
		case TAGF_A0.code:
			result[TAGF_A0] = tlv
		case TAGF_A1.code:
			result[TAGF_A1] = tlv
		case TAGF_A2.code:
			result[TAGF_A2] = tlv
		case TAGF_A5.code:
			result[TAGF_A5] = tlv
		case TAGF_AB.code:
			result[TAGF_AB] = tlv
		case TAGF_AC.code:
			result[TAGF_AC] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserFCP) parse6F(bytes []byte) (map[TagFCP]ber.TLV, error) {
	result := make(map[TagFCP]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGF_62.code:
			a, err := p.parse62(tlv.Value)
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

func (p parserFCP) parseFCP(bytes []byte) (map[TagFCP]ber.TLV, error) {
	result := make(map[TagFCP]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}

	tlv, err := bp.ParseTLV()
	if err != nil {
		return nil, err
	}

	switch t := bytesToHex(tlv.Tag); t {
	case TAGF_6F.code:
		a, err := p.parse6F(tlv.Value)
		if err != nil {
			return nil, err
		}
		for k, v := range a {
			result[k] = v
		}
	default:
		logUnknownTag(tlv, bytes)
	}

	return result, nil
}
