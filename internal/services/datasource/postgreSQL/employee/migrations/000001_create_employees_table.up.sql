CREATE TABLE IF NOT EXISTS employees
(
    id           BIGSERIAL    NOT NULL,
    employee_id  uuid         NOT NULL UNIQUE,
    company_id   uuid         NOT NULL,
    name         VARCHAR(200) NOT NULL,
    job_position job_position NOT NULL,
    compensation text         NOT NULL
);

CREATE TYPE job_position as ENUM
    (
        'programmer',
        'designer',
        'admin'
        );
