// Code generated by sqlc. DO NOT EDIT.

package company

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCompanyStmt, err = db.PrepareContext(ctx, createCompany); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCompany: %w", err)
	}
	if q.createCompanyAddressStmt, err = db.PrepareContext(ctx, createCompanyAddress); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCompanyAddress: %w", err)
	}
	if q.deleteCompanyStmt, err = db.PrepareContext(ctx, deleteCompany); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCompany: %w", err)
	}
	if q.getCompanyByIDStmt, err = db.PrepareContext(ctx, getCompanyByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetCompanyByID: %w", err)
	}
	if q.getCompanyDetailsByIDStmt, err = db.PrepareContext(ctx, getCompanyDetailsByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetCompanyDetailsByID: %w", err)
	}
	if q.listCompaniesStmt, err = db.PrepareContext(ctx, listCompanies); err != nil {
		return nil, fmt.Errorf("error preparing query ListCompanies: %w", err)
	}
	if q.listCompaniesDetailsStmt, err = db.PrepareContext(ctx, listCompaniesDetails); err != nil {
		return nil, fmt.Errorf("error preparing query ListCompaniesDetails: %w", err)
	}
	if q.updateCompanyStmt, err = db.PrepareContext(ctx, updateCompany); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCompany: %w", err)
	}
	if q.updateCompanyAddressStmt, err = db.PrepareContext(ctx, updateCompanyAddress); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCompanyAddress: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCompanyStmt != nil {
		if cerr := q.createCompanyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCompanyStmt: %w", cerr)
		}
	}
	if q.createCompanyAddressStmt != nil {
		if cerr := q.createCompanyAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCompanyAddressStmt: %w", cerr)
		}
	}
	if q.deleteCompanyStmt != nil {
		if cerr := q.deleteCompanyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCompanyStmt: %w", cerr)
		}
	}
	if q.getCompanyByIDStmt != nil {
		if cerr := q.getCompanyByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCompanyByIDStmt: %w", cerr)
		}
	}
	if q.getCompanyDetailsByIDStmt != nil {
		if cerr := q.getCompanyDetailsByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCompanyDetailsByIDStmt: %w", cerr)
		}
	}
	if q.listCompaniesStmt != nil {
		if cerr := q.listCompaniesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCompaniesStmt: %w", cerr)
		}
	}
	if q.listCompaniesDetailsStmt != nil {
		if cerr := q.listCompaniesDetailsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCompaniesDetailsStmt: %w", cerr)
		}
	}
	if q.updateCompanyStmt != nil {
		if cerr := q.updateCompanyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCompanyStmt: %w", cerr)
		}
	}
	if q.updateCompanyAddressStmt != nil {
		if cerr := q.updateCompanyAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCompanyAddressStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                        DBTX
	tx                        *sql.Tx
	createCompanyStmt         *sql.Stmt
	createCompanyAddressStmt  *sql.Stmt
	deleteCompanyStmt         *sql.Stmt
	getCompanyByIDStmt        *sql.Stmt
	getCompanyDetailsByIDStmt *sql.Stmt
	listCompaniesStmt         *sql.Stmt
	listCompaniesDetailsStmt  *sql.Stmt
	updateCompanyStmt         *sql.Stmt
	updateCompanyAddressStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                        tx,
		tx:                        tx,
		createCompanyStmt:         q.createCompanyStmt,
		createCompanyAddressStmt:  q.createCompanyAddressStmt,
		deleteCompanyStmt:         q.deleteCompanyStmt,
		getCompanyByIDStmt:        q.getCompanyByIDStmt,
		getCompanyDetailsByIDStmt: q.getCompanyDetailsByIDStmt,
		listCompaniesStmt:         q.listCompaniesStmt,
		listCompaniesDetailsStmt:  q.listCompaniesDetailsStmt,
		updateCompanyStmt:         q.updateCompanyStmt,
		updateCompanyAddressStmt:  q.updateCompanyAddressStmt,
	}
}