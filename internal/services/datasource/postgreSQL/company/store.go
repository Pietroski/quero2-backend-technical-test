package company

import "database/sql"

// Store provides all functions to execute db queries
type Store interface {
	Querier
}

type companyStore struct {
	*Queries
	db *sql.DB
}

// NewStore instantiates a company store object returning the store interface.
func NewStore(db *sql.DB) Store {
	negStore := &companyStore{
		Queries: New(db),
		db:      db,
	}

	return negStore
}
