-- name: GetEmployee :one
SELECT *
FROM employees
WHERE employee_id = $1
LIMIT 1;

-- name: ListAllEmployees :many
SELECT *
FROM employees
ORDER BY name;

-- name: ListPaginatedEmployees :many
SELECT *
FROM employees
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CreateEmployee :one
INSERT INTO employees (employee_id, name, job_position, compensation)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteEmployee :exec
DELETE
FROM employees
WHERE employee_id = $1;
