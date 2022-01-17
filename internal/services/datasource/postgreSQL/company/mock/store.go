// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company (interfaces: Store)

// Package mockdbcompany is a generated GoMock package.
package mockdbcompany

import (
	context "context"
	reflect "reflect"

	company "github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockStore) CreateCompany(arg0 context.Context, arg1 company.CreateCompanyParams) (company.Companies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", arg0, arg1)
	ret0, _ := ret[0].(company.Companies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockStoreMockRecorder) CreateCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockStore)(nil).CreateCompany), arg0, arg1)
}

// CreateCompanyAddress mocks base method.
func (m *MockStore) CreateCompanyAddress(arg0 context.Context, arg1 company.CreateCompanyAddressParams) (company.CompanyAddresses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompanyAddress", arg0, arg1)
	ret0, _ := ret[0].(company.CompanyAddresses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompanyAddress indicates an expected call of CreateCompanyAddress.
func (mr *MockStoreMockRecorder) CreateCompanyAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompanyAddress", reflect.TypeOf((*MockStore)(nil).CreateCompanyAddress), arg0, arg1)
}

// DeleteCompany mocks base method.
func (m *MockStore) DeleteCompany(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockStoreMockRecorder) DeleteCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockStore)(nil).DeleteCompany), arg0, arg1)
}

// GetCompanyByID mocks base method.
func (m *MockStore) GetCompanyByID(arg0 context.Context, arg1 uuid.UUID) (company.Companies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyByID", arg0, arg1)
	ret0, _ := ret[0].(company.Companies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyByID indicates an expected call of GetCompanyByID.
func (mr *MockStoreMockRecorder) GetCompanyByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyByID", reflect.TypeOf((*MockStore)(nil).GetCompanyByID), arg0, arg1)
}

// GetCompanyDetailsByID mocks base method.
func (m *MockStore) GetCompanyDetailsByID(arg0 context.Context, arg1 uuid.UUID) (company.GetCompanyDetailsByIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyDetailsByID", arg0, arg1)
	ret0, _ := ret[0].(company.GetCompanyDetailsByIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyDetailsByID indicates an expected call of GetCompanyDetailsByID.
func (mr *MockStoreMockRecorder) GetCompanyDetailsByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyDetailsByID", reflect.TypeOf((*MockStore)(nil).GetCompanyDetailsByID), arg0, arg1)
}

// ListCompanies mocks base method.
func (m *MockStore) ListCompanies(arg0 context.Context, arg1 company.ListCompaniesParams) ([]company.Companies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCompanies", arg0, arg1)
	ret0, _ := ret[0].([]company.Companies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCompanies indicates an expected call of ListCompanies.
func (mr *MockStoreMockRecorder) ListCompanies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCompanies", reflect.TypeOf((*MockStore)(nil).ListCompanies), arg0, arg1)
}

// ListCompaniesDetails mocks base method.
func (m *MockStore) ListCompaniesDetails(arg0 context.Context, arg1 company.ListCompaniesDetailsParams) ([]company.ListCompaniesDetailsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCompaniesDetails", arg0, arg1)
	ret0, _ := ret[0].([]company.ListCompaniesDetailsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCompaniesDetails indicates an expected call of ListCompaniesDetails.
func (mr *MockStoreMockRecorder) ListCompaniesDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCompaniesDetails", reflect.TypeOf((*MockStore)(nil).ListCompaniesDetails), arg0, arg1)
}

// UpdateCompany mocks base method.
func (m *MockStore) UpdateCompany(arg0 context.Context, arg1 company.UpdateCompanyParams) (company.Companies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", arg0, arg1)
	ret0, _ := ret[0].(company.Companies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockStoreMockRecorder) UpdateCompany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockStore)(nil).UpdateCompany), arg0, arg1)
}

// UpdateCompanyAddress mocks base method.
func (m *MockStore) UpdateCompanyAddress(arg0 context.Context, arg1 company.UpdateCompanyAddressParams) (company.CompanyAddresses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompanyAddress", arg0, arg1)
	ret0, _ := ret[0].(company.CompanyAddresses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCompanyAddress indicates an expected call of UpdateCompanyAddress.
func (mr *MockStoreMockRecorder) UpdateCompanyAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompanyAddress", reflect.TypeOf((*MockStore)(nil).UpdateCompanyAddress), arg0, arg1)
}