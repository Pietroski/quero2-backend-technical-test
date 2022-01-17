#!/usr/bin/env bash

migrate create -ext sql -dir internal/services/datasource/postgreSQL/company/migrations -seq create_companies_table
