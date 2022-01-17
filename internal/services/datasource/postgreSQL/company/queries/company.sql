-- name: CreateCompany :one
INSERT INTO companies (company_id,
                       company_address_id,
                       name,
                       phone_number)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateCompanyAddress :one
INSERT INTO company_addresses (company_address_id,
                               postal_code,
                               street,
                               number,
                               address_complement,
                               neighbourhood,
                               city,
                               state)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateCompany :one
UPDATE companies
SET (company_address_id,
     name,
     phone_number) = (
                      $2, $3, $4
    )
WHERE company_id = $1
RETURNING *;

-- name: UpdateCompanyAddress :one
UPDATE company_addresses
SET postal_code        = $2,
    street             = $3,
    number             = $4,
    address_complement = $5,
    neighbourhood      = $6,
    city               = $7,
    state              = $8
WHERE company_address_id = $1
RETURNING *;

-- name: GetCompanyByID :one
SELECT *
FROM companies
WHERE company_id = $1
LIMIT 1;

-- name: GetCompanyDetailsByID :one
SELECT *
FROM companies c
         INNER JOIN company_addresses ca
                    ON c.company_address_id = ca.company_address_id
WHERE c.company_id = $1
LIMIT 1;

-- name: ListCompanies :many
SELECT *
FROM companies
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: ListCompaniesDetails :many
SELECT *
FROM companies c
         INNER JOIN company_addresses ca
                    ON c.company_address_id = ca.company_address_id
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: DeleteCompany :exec
DELETE
FROM companies
WHERE company_id = $1;
