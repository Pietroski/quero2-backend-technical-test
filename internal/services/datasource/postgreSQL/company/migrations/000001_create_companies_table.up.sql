CREATE TABLE IF NOT EXISTS companies
(
    id                 BIGSERIAL    NOT NULL,
    company_id         uuid         NOT NULL UNIQUE,
    company_address_id uuid         NOT NULL,
    name               VARCHAR(255) NOT NULL,
    phone_number       VARCHAR(11)  NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (company_address_id) REFERENCES company_addresses (company_address_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS company_addresses
(
    id                 BIGSERIAL NOT NULL,
    company_address_id uuid      NOT NULL UNIQUE,
    postal_code        TEXT      NOT NULL,
    street             TEXT      NOT NULL,
    number             TEXT      NOT NULL,
    address_complement TEXT      NOT NULL,
    neighbourhood      TEXT      NOT NULL,
    city               TEXT      NOT NULL,
    state              TEXT      NOT NULL,
    PRIMARY KEY (id)
)
