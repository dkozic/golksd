package sd

import "github.com/dkozic/golksd/ber"

type parserB struct{}

func newParserB() parserB {
	p := parserB{}
	return p
}

func (p parserB) parseA7(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_A7_83.code:
			result[TAGB_A7_83] = tlv
		case TAGB_A7_84.code:
			result[TAGB_A7_84] = tlv
		case TAGB_A7_85.code:
			result[TAGB_A7_85] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseA8(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_A8_83.code:
			result[TAGB_A8_83] = tlv
		case TAGB_A8_84.code:
			result[TAGB_A8_84] = tlv
		case TAGB_A8_85.code:
			result[TAGB_A8_85] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseA9(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_A9_83.code:
			result[TAGB_A9_83] = tlv
		case TAGB_A9_84.code:
			result[TAGB_A9_84] = tlv
		case TAGB_A9_85.code:
			result[TAGB_A9_85] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseA1(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_A7.code:
			a, err := p.parseA7(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_A8.code:
			a, err := p.parseA8(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_A9.code:
			a, err := p.parseA9(tlv.Value)
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

func (p parserB) parseA4(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_96.code:
			result[TAGB_96] = tlv
		case TAGB_97.code:
			result[TAGB_97] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseAD(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_9F1F.code:
			result[TAGB_9F1F] = tlv
		case TAGB_9F20.code:
			result[TAGB_9F20] = tlv
		case TAGB_9F21.code:
			result[TAGB_9F21] = tlv
		case TAGB_9F22.code:
			result[TAGB_9F22] = tlv
		case TAGB_9F23.code:
			result[TAGB_9F23] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseAE(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_98.code:
			result[TAGB_98] = tlv
		case TAGB_9C.code:
			result[TAGB_9C] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseA5(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_9D.code:
			result[TAGB_9D] = tlv
		case TAGB_9E.code:
			result[TAGB_9E] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseAF(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_9F26.code:
			result[TAGB_9F26] = tlv
		case TAGB_9F27.code:
			result[TAGB_9F27] = tlv
		case TAGB_9F28.code:
			result[TAGB_9F28] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseB0(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_9F29.code:
			result[TAGB_9F29] = tlv
		case TAGB_9F2A.code:
			result[TAGB_9F2A] = tlv
		case TAGB_9F2B.code:
			result[TAGB_9F2B] = tlv
		case TAGB_9F2C.code:
			result[TAGB_9F2C] = tlv
		case TAGB_9F2D.code:
			result[TAGB_9F2D] = tlv
		case TAGB_9F2E.code:
			result[TAGB_9F2E] = tlv
		case TAGB_9F2F.code:
			result[TAGB_9F2F] = tlv
		case TAGB_9F30.code:
			result[TAGB_9F30] = tlv
		case TAGB_9F31.code:
			result[TAGB_9F31] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parse78(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_4F.code:
			result[TAGB_4F] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parse72(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for bp.HasNext() {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}
		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_80.code:
			result[TAGB_80] = tlv
		case TAGB_A1.code:
			a, err := p.parseA1(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_A4.code:
			a, err := p.parseA4(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_98.code:
			result[TAGB_98] = tlv
		case TAGB_99.code:
			result[TAGB_99] = tlv
		case TAGB_9A.code:
			result[TAGB_9A] = tlv
		case TAGB_AD.code:
			a, err := p.parseAD(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_AE.code:
			a, err := p.parseAE(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_A5.code:
			a, err := p.parseA5(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_9F24.code:
			result[TAGB_9F24] = tlv
		case TAGB_9F25.code:
			result[TAGB_9F25] = tlv
		case TAGB_AF.code:
			a, err := p.parseAF(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_B0.code:
			a, err := p.parseB0(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_9F32.code:
			result[TAGB_9F32] = tlv
		default:
			logUnknownTag(tlv, bytes)
		}
	}
	return result, nil
}

func (p parserB) parseFile(bytes []byte) (map[TagB]ber.TLV, error) {
	result := make(map[TagB]ber.TLV)
	bp := ber.BERParser{Bytes: bytes}
	for i := 0; i < 2; i++ {
		tlv, err := bp.ParseTLV()
		if err != nil {
			return nil, err
		}

		switch t := bytesToHex(tlv.Tag); t {
		case TAGB_78.code:
			a, err := p.parse78(tlv.Value)
			if err != nil {
				return nil, err
			}
			for k, v := range a {
				result[k] = v
			}
		case TAGB_72.code:
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
