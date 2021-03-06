// Code generated by sqlc. DO NOT EDIT.

package employee

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employees, error)
	DeleteEmployee(ctx context.Context, employeeID uuid.UUID) error
	GetEmployee(ctx context.Context, employeeID uuid.UUID) (Employees, error)
	ListAllEmployees(ctx context.Context) ([]Employees, error)
	ListPaginatedEmployees(ctx context.Context, arg ListPaginatedEmployeesParams) ([]Employees, error)
}

var _ Querier = (*Queries)(nil)
