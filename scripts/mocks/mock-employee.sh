#!/usr/bin/env bash

mockgen -package mockdbemployee \
        -destination internal/services/datasource/postgreSQL/employee/mock/store.go \
        github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/employee Store
