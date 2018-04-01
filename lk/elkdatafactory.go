package lk

import "github.com/dkozic/golksd/ber"

type dataFactory struct {
	fieldsDocument  map[TagDocument]ber.TLV
	fieldsPersonal  map[TagPersonal]ber.TLV
	fieldsResidence map[TagResidence]ber.TLV
	portrait        []byte
}

func newDataFactory(fieldsDocument map[TagDocument]ber.TLV, fieldsPersonal map[TagPersonal]ber.TLV, fieldsResidence map[TagResidence]ber.TLV, portrait []byte) dataFactory {
	f := dataFactory{fieldsDocument, fieldsPersonal, fieldsResidence, portrait}
	return f
}

func (f dataFactory) getDocument(tag TagDocument) string {
	tlv, ok := f.fieldsDocument[tag]
	if !ok {
		return ""
	}
	return tag.tagCharset.BytesToCharsetString(tlv.Value)

}

func (f dataFactory) getPersonal(tag TagPersonal) string {
	tlv, ok := f.fieldsPersonal[tag]
	if !ok {
		return ""
	}
	return tag.tagCharset.BytesToCharsetString(tlv.Value)

}

func (f dataFactory) getResidence(tag TagResidence) string {
	tlv, ok := f.fieldsResidence[tag]
	if !ok {
		return ""
	}
	return tag.tagCharset.BytesToCharsetString(tlv.Value)

}

// Document data
func (f dataFactory) docRegNo() string {
	return f.getDocument(TAGD_0A)
}

func (f dataFactory) issuingDate() string {
	return f.getDocument(TAGD_0D)
}

func (f dataFactory) expiryDate() string {
	return f.getDocument(TAGD_0E)
}

func (f dataFactory) issuingAuthority() string {
	return f.getDocument(TAGD_0F)
}

// personal data
func (f dataFactory) personalNumber() string {
	return f.getPersonal(TAGP_16)
}

func (f dataFactory) surname() string {
	return f.getPersonal(TAGP_17)
}

func (f dataFactory) givenName() string {
	return f.getPersonal(TAGP_18)
}

func (f dataFactory) parentGivenName() string {
	return f.getPersonal(TAGP_19)
}

func (f dataFactory) sex() string {
	return f.getPersonal(TAGP_1A)
}

func (f dataFactory) placeOfBirth() string {
	return f.getPersonal(TAGP_1B)
}

func (f dataFactory) stateOfBirth() string {
	return f.getPersonal(TAGP_1D)
}

func (f dataFactory) dateOfBirth() string {
	return f.getPersonal(TAGP_1E)
}

// residence data
func (f dataFactory) state() string {
	return f.getResidence(TAGR_20)
}

func (f dataFactory) comunity() string {
	return f.getResidence(TAGR_21)
}

func (f dataFactory) place() string {
	return f.getResidence(TAGR_22)
}

func (f dataFactory) street() string {
	return f.getResidence(TAGR_23)
}

func (f dataFactory) houseNumber() string {
	return f.getResidence(TAGR_24)
}

func (f dataFactory) apartmentNumber() string {
	return f.getResidence(TAGR_2A)
}

func (f dataFactory) create() ElkData {
	d := ElkData{}

	d.IssuingDate = f.issuingDate()
	d.ExpiryDate = f.expiryDate()
	d.IssuingAuthority = f.issuingAuthority()
	d.PersonalNumber = f.personalNumber()
	d.Surname = f.surname()
	d.GivenName = f.givenName()
	d.ParentGivenName = f.parentGivenName()
	d.Sex = f.sex()
	d.PlaceOfBirth = f.placeOfBirth()
	d.StateOfBirth = f.stateOfBirth()
	d.DateOfBirth = f.dateOfBirth()
	d.State = f.state()
	d.Community = f.comunity()
	d.Place = f.place()
	d.Street = f.street()
	d.HouseNumber = f.houseNumber()
	d.ApartmentNumber = f.apartmentNumber()
	d.Portrait = f.portrait

	return d
}
