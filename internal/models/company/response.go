package companyModel

import "github.com/google/uuid"

type CompanyResponse struct {
	ID               int64     `json:"-"`
	CompanyID        uuid.UUID `json:"companyID"`
	CompanyAddressID uuid.UUID `json:"companyAddressID"`
	Name             string    `json:"name"`
	PhoneNumber      string    `json:"phoneNumber"`
}

type CompanyAddressResponse struct {
	ID                int64     `json:"-"`
	CompanyAddressID  uuid.UUID `json:"companyAddressID"`
	PostalCode        string    `json:"postalCode"`
	Street            string    `json:"street"`
	Number            string    `json:"number"`
	AddressComplement string    `json:"addressComplement"`
	Neighbourhood     string    `json:"neighbourhood"`
	City              string    `json:"city"`
	State             string    `json:"state"`
}
