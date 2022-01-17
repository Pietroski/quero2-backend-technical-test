package companyModel

type CompanyRequest struct {
	Name                    string                  `json:"name" binding:"required"`
	PhoneNumber             string                  `json:"phoneNumber" binding:"required"`
	CompanyAddressesRequest CompanyAddressesRequest `json:"companyAddressesRequest" binding:"required"`
}

type CompanyAddressesRequest struct {
	PostalCode        string `json:"postalCode" binding:"required"`
	Street            string `json:"street" binding:"required"`
	Number            string `json:"number" binding:"required"`
	AddressComplement string `json:"addressComplement" binding:"required"`
	Neighbourhood     string `json:"neighbourhood" binding:"required"`
	City              string `json:"city" binding:"required"`
	State             string `json:"state" binding:"required"`
}
