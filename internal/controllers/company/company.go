package companyController

import (
	companyModel "github.com/Pietroski/quero2-backend-technical-test/internal/models/company"
	"github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company"
	"github.com/Pietroski/quero2-backend-technical-test/internal/util/notification"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type (
	Server struct {
		store company.Store
	}
)

func NewServer(store company.Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) CreateCompanyWithAddress(ctx *gin.Context) {
	var (
		req companyModel.CompanyRequest
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	createdAddress, err := s.createAddress(ctx, req.CompanyAddressesRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	companyID := uuid.New()
	createCompanyPayload := company.CreateCompanyParams{
		CompanyID:        companyID,
		CompanyAddressID: createdAddress.CompanyAddressID,
		Name:             req.Name,
		PhoneNumber:      req.PhoneNumber,
	}

	createdCompany, err := s.store.CreateCompany(ctx, createCompanyPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	ccr := companyModel.CompanyResponse(createdCompany)

	ctx.JSON(http.StatusCreated, ccr)
}

func (s *Server) CreateAddress(ctx *gin.Context) {
	var req companyModel.CompanyAddressesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	createdAddress, err := s.createAddress(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	ctx.JSON(http.StatusCreated, createdAddress)
}

func (s *Server) createAddress(
	ctx *gin.Context,
	req companyModel.CompanyAddressesRequest,
) (companyModel.CompanyAddressResponse, error) {
	companyAddressID := uuid.New()
	companyAddressPayload := company.CreateCompanyAddressParams{
		CompanyAddressID:  companyAddressID,
		PostalCode:        req.PostalCode,
		Street:            req.Street,
		Number:            req.Number,
		AddressComplement: req.AddressComplement,
		Neighbourhood:     req.Neighbourhood,
		City:              req.City,
		State:             req.State,
	}

	companyAddress, err := s.store.CreateCompanyAddress(ctx, companyAddressPayload)
	if err != nil {
		return companyModel.CompanyAddressResponse{}, err
	}

	ca := companyModel.CompanyAddressResponse(companyAddress)

	return ca, nil
}

func (s *Server) UpdateCompanyWithAddress(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) UpdateCompanyAddress(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) updateCompanyAddress(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) GetCompanyWithAddress(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) GetCompanyAddress(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) DeleteCompanyWithAddress(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) DeleteCompanyAddress(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) ListCompaniesWithAddresses(ctx *gin.Context) {
	// TODO: implement me
}

func (s *Server) ListCompaniesAddresses(ctx *gin.Context) {
	// TODO: implement me
}
