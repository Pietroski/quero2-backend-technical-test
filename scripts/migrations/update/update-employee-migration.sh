#!/usr/bin/env bash

migration_name=${1?"migration name required"}

migrate create -ext sql -dir internal/services/datasource/postgreSQL/employee/migrations -seq $migration_name
