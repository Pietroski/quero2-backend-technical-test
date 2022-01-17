#!/usr/bin/env bash

mockgen -package mockdbcompany \
        -destination internal/services/datasource/postgreSQL/company/mock/store.go \
        github.com/Pietroski/quero2-backend-technical-test/internal/services/datasource/postgreSQL/company Store
