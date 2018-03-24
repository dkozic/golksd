package sd

import "github.com/dkozic/golksd/ber"

type parserA struct{}

func newParserA() parserA {
	p := parserA{}
	return p
}

func (p parserA) parseA2(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_A2_83.code:
			result[TAGA_A2_83] = tlv
		case TAGA_A2_84.code:
			result[TAGA_A2_84] = tlv
		case TAGA_A2_85.code:
			result[TAGA_A2_85] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserA) parseA1(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_A2.code:
			a2, err := p.parseA2(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a2 {
				result[k] = v
			}
		case TAGA_86.code:
			result[TAGA_86] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserA) parseA3(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_87.code:
			result[TAGA_87] = tlv
		case TAGA_88.code:
			result[TAGA_88] = tlv
		case TAGA_89.code:
			result[TAGA_89] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserA) parseA4(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_8B.code:
			result[TAGA_8B] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserA) parseA5(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_90.code:
			result[TAGA_90] = tlv
		case TAGA_91.code:
			result[TAGA_91] = tlv
		case TAGA_92.code:
			result[TAGA_92] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserA) parseA6(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_94.code:
			result[TAGA_94] = tlv
		case TAGA_95.code:
			result[TAGA_95] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserA) parse78(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_4F.code:
			result[TAGA_4F] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserA) parse71(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_80.code:
			result[TAGA_80] = tlv
		case TAGA_9F33.code:
			result[TAGA_9F33] = tlv
		case TAGA_9F34.code:
			result[TAGA_9F34] = tlv
		case TAGA_9F35.code:
			result[TAGA_9F35] = tlv
		case TAGA_9F36.code:
			result[TAGA_9F36] = tlv
		case TAGA_9F37.code:
			result[TAGA_9F37] = tlv
		case TAGA_9F38.code:
			result[TAGA_9F38] = tlv
		case TAGA_81.code:
			result[TAGA_81] = tlv
		case TAGA_82.code:
			result[TAGA_82] = tlv
		case TAGA_A1.code:
			a1, err := p.parseA1(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a1 {
				result[k] = v
			}
		case TAGA_A3.code:
			a3, err := p.parseA3(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a3 {
				result[k] = v
			}
		case TAGA_8A.code:
			result[TAGA_8A] = tlv
		case TAGA_A4.code:
			a4, err := p.parseA4(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a4 {
				result[k] = v
			}
		case TAGA_8C.code:
			result[TAGA_8C] = tlv
		case TAGA_8D.code:
			result[TAGA_8D] = tlv
		case TAGA_8E.code:
			result[TAGA_8E] = tlv
		case TAGA_8F.code:
			result[TAGA_8F] = tlv
		case TAGA_A5.code:
			a5, err := p.parseA5(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a5 {
				result[k] = v
			}
		case TAGA_93.code:
			result[TAGA_93] = tlv
		case TAGA_A6.code:
			a6, err := p.parseA6(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a6 {
				result[k] = v
			}
		default:
			logUnknownTag(tlv, bytes)
		}

	}
	return result, nil
}

func (p parserA) parseFile(bytes []byte) (map[TagA]ber.TLV, error) {
	result := make(map[TagA]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for i := 0; i < 2; i++ {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGA_78.code:
			a, err := p.parse78(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGA_71.code:
			a, err := p.parse71(tlv.Value)
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
