package companyController_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	companyFactory "github.com/Pietroski/quero2-backend-technical-test/internal/factories/company"
	mockdbcompany "github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_CreateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdbcompany.NewMockStore(ctrl)
	server := companyFactory.NewCompanyFactory(store)

	addressPar := createRandomAddressParams()
	addressResult := createRandomAddressResult(1)
	compPar := createRandomCompanyParams(addressResult.CompanyAddressID)
	compResult := createRandomCompanyResult(1, addressResult.CompanyAddressID)

	{
		store.
			EXPECT().
			CreateCompanyAddress(gomock.Any(), EqCreateCompanyAddressParams(*addressPar)).
			Times(1).
			Return(*addressResult, nil)

		store.
			EXPECT().
			CreateCompany(gomock.Any(), EqCreateCompanyParams(*compPar)).
			Times(1).
			Return(*compResult, nil)
	}

	cr := createRandomCompReq(*createRandomCompAddressReq())

	postReq, err := json.Marshal(cr)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/companies")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(postReq))
	require.NoError(t, err)
	server.Router.ServeHTTP(recorder, req)

	// check response
	require.Equal(t, http.StatusCreated, recorder.Code)
}
