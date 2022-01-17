package companyFactory

import (
	"github.com/Pietroski/quero2-backend-technical-test/internal/controllers/company"
	"github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type (
	assembler struct {
		store    company.Store
		Router   *gin.Engine
		handlers *companyController.Server
	}
)

func NewCompanyFactory(store company.Store) *assembler {
	a := &assembler{
		store:    store,
		handlers: companyController.NewServer(store),
	}
	a.handler()

	return a
}

func (a *assembler) handler() {
	gin.ForceConsoleColor()
	router := gin.New()

	// CORS middleware - default
	router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		v1.POST("/companies", a.handlers.CreateCompanyWithAddress)
	}

	a.Router = router
}

func (a *assembler) Start(address string) error {
	return a.Router.Run(address)
}
