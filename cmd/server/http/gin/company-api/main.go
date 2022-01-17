package main

import (
	"database/sql"
	companyFactory "github.com/Pietroski/quero2-backend-technical-test/internal/factories/company"
	"github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company"
	"log"
)

var (
	// apiConfig models.APIStruct
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("apiConfig.DBDriver", "apiConfig.DBDataSourceName")
	if err != nil {
		log.Fatalf("Error connecting to database -> %s", err.Error())
	}
}

func main() {
	companyStore := company.NewStore(dbConn)

	companyServer := companyFactory.NewCompanyFactory(companyStore)
	if err = companyServer.Start("server-address"); err != nil {
		log.Fatalf("cannot start server -> %s", err.Error())
	}
}
