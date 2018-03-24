package sd

type EsdData struct {
	// SD_DOCUMENT_DATA
	StateIssuing       string
	CompetentAuthority string
	AuthorityIssuing   string
	UnambiguousNumber  string
	IssuingDate        string
	ExpiryDate         string
	SerialNumber       string

	// SD_VEHICLE_DATA
	DateOfFirstRegistration     string
	YearOfProduction            string
	VehicleMake                 string
	VehicleType                 string
	CommercialDescription       string
	VehicleIDNumber             string
	RegistrationNumberOfVehicle string
	MaximumNetPower             string
	EngineCapacity              string
	TypeOfFuel                  string
	PowerWeightRatio            string
	VehicleMass                 string
	MaximumPermissibleLadenMass string
	TypeApprovalNumber          string
	NumberOfSeats               string
	NumberOfStandingPlaces      string
	EngineIDNumber              string
	NumberOfAxles               string
	VehicleCategory             string
	ColourOfVehicle             string
	RestrictionToChangeOwner    string
	VehicleLoad                 string

	// SD_PERSONAL_DATA
	HolderSurnameOrBusinessName string
	HolderName                  string
	HolderAddress               string
	OwnerPersonalNo             string
	OwnerSurnameOrBusinessName  string
	OwnerName                   string
	OwnerAddress                string
	UserPersonalNo              string
	UserSurnameOrBusinessName   string
	UserName                    string
	UserAddress                 string
}
