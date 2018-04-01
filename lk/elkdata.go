package lk

type ElkData struct {

	// Document data
	DocRegNo         string
	IssuingDate      string
	ExpiryDate       string
	IssuingAuthority string

	// Fixed personal data
	PersonalNumber   string
	Surname          string
	GivenName        string
	ParentGivenName  string
	Sex              string
	PlaceOfBirth     string
	StateOfBirth     string
	DateOfBirth      string
	CommunityOfBirth string

	// Variable personal data
	State           string
	Community       string
	Place           string
	Street          string
	HouseNumber     string
	HouseLetter     string
	Entrance        string
	Floor           string
	ApartmentNumber string

	// Portrait
	Portrait []byte
}
