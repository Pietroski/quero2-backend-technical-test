package companyController_test

import (
	companyModel "github.com/Pietroski/quero2-backend-technical-test/internal/models/company"
	"github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company"
	"github.com/google/uuid"
)

func createRandomAddressParams() *company.CreateCompanyAddressParams {
	car := &company.CreateCompanyAddressParams{
		PostalCode:        "postal-code",
		Street:            "street",
		Number:            "number",
		AddressComplement: "address-complement",
		Neighbourhood:     "neighbourhood",
		City:              "city",
		State:             "state",
	}

	return car
}

func createRandomAddressResult(id int64) *company.CompanyAddresses {
	car := &company.CompanyAddresses{
		ID:                id,
		CompanyAddressID:  uuid.UUID{},
		PostalCode:        "postal-code",
		Street:            "street",
		Number:            "number",
		AddressComplement: "address-complement",
		Neighbourhood:     "neighbourhood",
		City:              "city",
		State:             "state",
	}

	return car
}

func createRandomCompanyParams(companyAddressID uuid.UUID) *company.CreateCompanyParams {
	cr := &company.CreateCompanyParams{
		CompanyID:        uuid.UUID{},
		CompanyAddressID: companyAddressID,
		Name:             "name",
		PhoneNumber:      "phone-number",
	}

	return cr
}

func createRandomCompanyResult(id int64, companyAddressID uuid.UUID) *company.Companies {
	cr := &company.Companies{
		ID:               id,
		CompanyID:        uuid.UUID{},
		CompanyAddressID: companyAddressID,
		Name:             "name",
		PhoneNumber:      "phone-number",
	}

	return cr
}

func createRandomCompReq(
	companyAddressesRequest companyModel.CompanyAddressesRequest,
) *companyModel.CompanyRequest {
	cr := &companyModel.CompanyRequest{
		Name:                    "name",
		PhoneNumber:             "phone-number",
		CompanyAddressesRequest: companyAddressesRequest,
	}

	return cr
}

func createRandomCompAddressReq() *companyModel.CompanyAddressesRequest {
	car := &companyModel.CompanyAddressesRequest{
		PostalCode:        "postal-code",
		Street:            "street",
		Number:            "number",
		AddressComplement: "address-complement",
		Neighbourhood:     "neighbourhood",
		City:              "city",
		State:             "state",
	}

	return car
}
