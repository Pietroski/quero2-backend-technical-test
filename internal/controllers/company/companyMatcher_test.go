package companyController_test

import (
	"fmt"
	"github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company"
	"github.com/golang/mock/gomock"
	//"github.com/google/uuid"
	"reflect"
)

func Eq(x interface{}) Matcher { return eqMatcher{x} }

type Matcher interface {
	// Matches returns whether x is a match.
	Matches(x interface{}) bool

	// String describes what the matcher matches.
	String() string
}

type eqMatcher struct {
	x interface{}
}

func (e eqMatcher) Matches(x interface{}) bool {
	return reflect.DeepEqual(e.x, x)
}

func (e eqMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.x)
}

type eqCreateCompanyAddressParamsMatcher struct {
	arg company.CreateCompanyAddressParams
}

type eqCreateCompanyParamsMatcher struct {
	arg company.CreateCompanyParams
}

func (e eqCreateCompanyAddressParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(company.CreateCompanyAddressParams)
	if !ok {
		return false
	}

	return reflect.DeepEqual(e.arg.Street, arg.Street) &&
		reflect.DeepEqual(e.arg.Number, arg.Number) &&
		reflect.DeepEqual(e.arg.AddressComplement, arg.AddressComplement) &&
		reflect.DeepEqual(e.arg.Neighbourhood, arg.Neighbourhood) &&
		reflect.DeepEqual(e.arg.City, arg.City) &&
		reflect.DeepEqual(e.arg.State, arg.State) &&
		reflect.DeepEqual(e.arg.PostalCode, arg.PostalCode)
}

func (e eqCreateCompanyAddressParamsMatcher) String() string {
	return fmt.Sprintf("does not match arg %v", e.arg)
}

func EqCreateCompanyAddressParams(arg company.CreateCompanyAddressParams) gomock.Matcher {
	return eqCreateCompanyAddressParamsMatcher{arg}
}

func (e eqCreateCompanyParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(company.CreateCompanyParams)
	if !ok {
		return false
	}

	return reflect.DeepEqual(e.arg.Name, arg.Name) &&
		reflect.DeepEqual(e.arg.PhoneNumber, arg.PhoneNumber) &&
		reflect.DeepEqual(e.arg.CompanyAddressID, arg.CompanyAddressID)
}

func (e eqCreateCompanyParamsMatcher) String() string {
	return fmt.Sprintf("does not match arg %v", e.arg)
}

func EqCreateCompanyParams(arg company.CreateCompanyParams) gomock.Matcher {
	return eqCreateCompanyParamsMatcher{arg}
}
