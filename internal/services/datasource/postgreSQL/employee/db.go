// Code generated by sqlc. DO NOT EDIT.

package employee

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
	if q.createEmployeeStmt, err = db.PrepareContext(ctx, createEmployee); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEmployee: %w", err)
	}
	if q.deleteEmployeeStmt, err = db.PrepareContext(ctx, deleteEmployee); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEmployee: %w", err)
	}
	if q.getEmployeeStmt, err = db.PrepareContext(ctx, getEmployee); err != nil {
		return nil, fmt.Errorf("error preparing query GetEmployee: %w", err)
	}
	if q.listAllEmployeesStmt, err = db.PrepareContext(ctx, listAllEmployees); err != nil {
		return nil, fmt.Errorf("error preparing query ListAllEmployees: %w", err)
	}
	if q.listPaginatedEmployeesStmt, err = db.PrepareContext(ctx, listPaginatedEmployees); err != nil {
		return nil, fmt.Errorf("error preparing query ListPaginatedEmployees: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createEmployeeStmt != nil {
		if cerr := q.createEmployeeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEmployeeStmt: %w", cerr)
		}
	}
	if q.deleteEmployeeStmt != nil {
		if cerr := q.deleteEmployeeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEmployeeStmt: %w", cerr)
		}
	}
	if q.getEmployeeStmt != nil {
		if cerr := q.getEmployeeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEmployeeStmt: %w", cerr)
		}
	}
	if q.listAllEmployeesStmt != nil {
		if cerr := q.listAllEmployeesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAllEmployeesStmt: %w", cerr)
		}
	}
	if q.listPaginatedEmployeesStmt != nil {
		if cerr := q.listPaginatedEmployeesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPaginatedEmployeesStmt: %w", cerr)
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
	db                         DBTX
	tx                         *sql.Tx
	createEmployeeStmt         *sql.Stmt
	deleteEmployeeStmt         *sql.Stmt
	getEmployeeStmt            *sql.Stmt
	listAllEmployeesStmt       *sql.Stmt
	listPaginatedEmployeesStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                         tx,
		tx:                         tx,
		createEmployeeStmt:         q.createEmployeeStmt,
		deleteEmployeeStmt:         q.deleteEmployeeStmt,
		getEmployeeStmt:            q.getEmployeeStmt,
		listAllEmployeesStmt:       q.listAllEmployeesStmt,
		listPaginatedEmployeesStmt: q.listPaginatedEmployeesStmt,
	}
}
