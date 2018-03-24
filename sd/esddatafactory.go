package sd

import "github.com/dkozic/golksd/ber"

type dataFactory struct {
	fieldsA map[TagA]ber.TLV
	fieldsB map[TagB]ber.TLV
	fieldsC map[TagC]ber.TLV
	fieldsD map[TagD]ber.TLV
}

func newDataFactory(fieldsA map[TagA]ber.TLV, fieldsB map[TagB]ber.TLV, fieldsC map[TagC]ber.TLV, fieldsD map[TagD]ber.TLV) dataFactory {
	f := dataFactory{fieldsA, fieldsB, fieldsC, fieldsD}
	return f
}

func (f dataFactory) getA(tag TagA) string {
	tlv, ok := f.fieldsA[tag]
	if !ok {
		return ""
	}
	return tag.tagCharset.BytesToCharsetString(tlv.Value)

}

func (f dataFactory) getB(tag TagB) string {
	tlv, ok := f.fieldsB[tag]
	if !ok {
		return ""
	}
	return tag.tagCharset.BytesToCharsetString(tlv.Value)

}

func (f dataFactory) getC(tag TagC) string {
	tlv, ok := f.fieldsC[tag]
	if !ok {
		return ""
	}
	return tag.tagCharset.BytesToCharsetString(tlv.Value)

}

func (f dataFactory) getD(tag TagD) string {
	tlv, ok := f.fieldsD[tag]
	if !ok {
		return ""
	}
	return tag.tagCharset.BytesToCharsetString(tlv.Value)

}

// SD_DOCUMENT_DATA
func (f dataFactory) stateIssuing() string {
	return f.getA(TAGA_9F33)
}

func (f dataFactory) competentAuthority() string {
	return f.getA(TAGA_9F35)
}

func (f dataFactory) authorityIssuing() string {
	return f.getA(TAGA_9F36)
}

func (f dataFactory) unambiguousNumber() string {
	return f.getA(TAGA_9F38)
}

func (f dataFactory) issuingDate() string {
	return f.getA(TAGA_8E)
}

func (f dataFactory) expiryDate() string {
	return f.getA(TAGA_8D)
}

func (f dataFactory) serialNumber() string {
	return f.getD(TAGD_C9)
}

// SD_VEHICLE_DATA
func (f dataFactory) dateOfFirstRegistration() string {
	return f.getA(TAGA_82)
}

func (f dataFactory) yearOfProduction() string {
	return f.getC(TAGC_C5)
}

func (f dataFactory) vehicleMake() string {
	return f.getA(TAGA_87)
}

func (f dataFactory) vehicleType() string {
	return f.getA(TAGA_88)
}

func (f dataFactory) commercialDescription() string {
	return f.getA(TAGA_89)
}

func (f dataFactory) vehicleIDNumber() string {
	return f.getA(TAGA_8A)
}

func (f dataFactory) registrationNumberOfVehicle() string {
	return f.getA(TAGA_81)
}

func (f dataFactory) maximumNetPower() string {
	return f.getA(TAGA_91)
}

func (f dataFactory) engineCapacity() string {
	return f.getA(TAGA_90)
}

func (f dataFactory) typeOfFuel() string {
	return f.getA(TAGA_92)
}

func (f dataFactory) powerWeightRatio() string {
	return f.getA(TAGA_93)
}

func (f dataFactory) vehicleMass() string {
	return f.getA(TAGA_8C)
}

func (f dataFactory) maximumPermissibleLadenMass() string {
	return f.getA(TAGA_8B)
}

func (f dataFactory) typeApprovalNumber() string {
	return f.getA(TAGA_8F)
}

func (f dataFactory) numberOfSeats() string {
	return f.getA(TAGA_94)
}

func (f dataFactory) numberOfStandingPlaces() string {
	return f.getA(TAGA_95)
}

func (f dataFactory) engineIDNumber() string {
	return f.getB(TAGB_9E)
}

func (f dataFactory) numberOfAxles() string {
	return f.getB(TAGB_99)
}

func (f dataFactory) vehicleCategory() string {
	return f.getB(TAGB_98)
}

func (f dataFactory) colourOfVehicle() string {
	return f.getB(TAGB_9F24)
}

func (f dataFactory) restrictionToChangeOwner() string {
	return f.getC(TAGC_C1)
}

func (f dataFactory) vehicleLoad() string {
	return f.getC(TAGC_C4)
}

// SD_PERSONAL_DATA
func (f dataFactory) holderSurnameOrBusinessName() string {
	return f.getA(TAGA_A2_83)
}

func (f dataFactory) holderName() string {
	return f.getA(TAGA_A2_84)
}

func (f dataFactory) holderAddress() string {
	return f.getA(TAGA_A2_85)
}

func (f dataFactory) ownerPersonalNo() string {
	return f.getC(TAGC_C2)
}

func (f dataFactory) ownerSurnameOrBusinessName() string {
	return f.getB(TAGB_A7_83)
}

func (f dataFactory) ownerName() string {
	return f.getB(TAGB_A7_84)
}

func (f dataFactory) ownerAddress() string {
	return f.getB(TAGB_A7_85)
}

func (f dataFactory) userPersonalNo() string {
	return f.getC(TAGC_C3)
}

func (f dataFactory) userSurnameOrBusinessName() string {
	return f.getB(TAGB_A9_83)
}

func (f dataFactory) userName() string {
	return f.getB(TAGB_A9_84)
}

func (f dataFactory) userAddress() string {
	return f.getB(TAGB_A9_85)
}
func (f dataFactory) create() EsdData {
	d := EsdData{}
	d.StateIssuing = f.stateIssuing()
	d.CompetentAuthority = f.competentAuthority()
	d.AuthorityIssuing = f.authorityIssuing()
	d.UnambiguousNumber = f.unambiguousNumber()

	d.IssuingDate = f.issuingDate()
	d.ExpiryDate = f.expiryDate()
	d.SerialNumber = f.serialNumber()
	d.DateOfFirstRegistration = f.dateOfFirstRegistration()
	d.YearOfProduction = f.yearOfProduction()
	d.VehicleMake = f.vehicleMake()
	d.VehicleType = f.vehicleType()
	d.CommercialDescription = f.commercialDescription()
	d.VehicleIDNumber = f.vehicleIDNumber()
	d.RegistrationNumberOfVehicle = f.registrationNumberOfVehicle()
	d.MaximumNetPower = f.maximumNetPower()
	d.EngineCapacity = f.engineCapacity()
	d.TypeOfFuel = f.typeOfFuel()
	d.PowerWeightRatio = f.powerWeightRatio()
	d.VehicleMass = f.vehicleMass()
	d.MaximumPermissibleLadenMass = f.maximumPermissibleLadenMass()
	d.TypeApprovalNumber = f.typeApprovalNumber()
	d.NumberOfSeats = f.numberOfSeats()
	d.NumberOfStandingPlaces = f.numberOfStandingPlaces()
	d.EngineIDNumber = f.engineIDNumber()
	d.NumberOfAxles = f.numberOfAxles()
	d.VehicleCategory = f.vehicleCategory()
	d.ColourOfVehicle = f.colourOfVehicle()
	d.RestrictionToChangeOwner = f.restrictionToChangeOwner()
	d.VehicleLoad = f.vehicleLoad()
	d.HolderSurnameOrBusinessName = f.holderSurnameOrBusinessName()
	d.HolderName = f.holderName()
	d.HolderAddress = f.holderAddress()
	d.OwnerPersonalNo = f.ownerPersonalNo()
	d.OwnerSurnameOrBusinessName = f.ownerSurnameOrBusinessName()
	d.OwnerName = f.ownerName()
	d.OwnerAddress = f.ownerAddress()
	d.UserPersonalNo = f.userPersonalNo()
	d.UserSurnameOrBusinessName = f.userSurnameOrBusinessName()
	d.UserName = f.userName()
	d.UserAddress = f.userAddress()

	return d
}
