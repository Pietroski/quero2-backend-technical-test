#!/usr/bin/env bash

migrate create -ext sql -dir internal/services/datasource/postgreSQL/employee/migrations -seq create_employees_table
