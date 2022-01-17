
sqlc-init-default:
	sqlc init

sqlc-generate-default:
	sqlc generate

mock-companies:
	scripts/mocks/mock-company.sh

mock-employees:
	scripts/mocks/mock-employee.sh

mock-all:
	@make \
	mock-companies \
	mock-employees;
