MAIN=cmd/balance/main.go
OUTPUT=balance

MODULE=github.com/tksasha/balance
MODFILE=-modfile go.tool.mod

.PHONY: default
default: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@go fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go tool -modfile go.tool.mod gofumpt -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@go tool -modfile go.tool.mod golangci-lint run

.PHONY: test
test:
	@echo "go test"
	@go test ./... -count=1

.PHONY: update
update:
	@go test ./... -update

.PHONY: air
air:
	@go tool -modfile go.tool.mod air

.PHONY: run
run:
	@echo "go run (without live reloading)"
	@go run github.com/tksasha/balance/cmd/balance

.PHONY: build
build:
	@echo "go build"
	@go build -o $(OUTPUT) $(MAIN)

.PHONY: clear
clear:
	@echo "go clear"
	@rm $(OUTPUT)

.PHONY: clean
clean: clear

.PHONY: gen
gen: wire mockgen

.PHONY: mockgen
mockgen:
	@go tool -modfile go.tool.mod mockgen \
		-source internal/app/cash/interfaces.go \
		-package mocks \
		-destination internal/app/cash/test/mocks/interfaces.mock.go
	@go tool -modfile go.tool.mod mockgen \
		-source internal/app/category/interfaces.go \
		-package mocks \
		-destination internal/app/category/test/mocks/interfaces.mock.go
	@go tool -modfile go.tool.mod mockgen \
		-source internal/app/item/interfaces.go \
		-package mocks \
		-destination internal/app/item/test/mocks/interfaces.mock.go
	@go tool -modfile go.tool.mod mockgen \
		-source internal/app/index/interfaces.go \
		-package mocks \
		-destination internal/app/index/test/mocks/interfaces.mock.go
	@go tool -modfile go.tool.mod mockgen \
		-source internal/backoffice/category/interfaces.go \
		-package mocks \
		-destination internal/backoffice/category/test/mocks/interfaces.mock.go
	@go tool -modfile go.tool.mod mockgen \
		-source internal/backoffice/cash/interfaces.go \
		-package mocks \
		-destination internal/backoffice/cash/test/mocks/interfaces.mock.go
	@go tool -modfile go.tool.mod mockgen \
		-source internal/app/balance/interfaces.go \
		-package mocks \
		-destination internal/app/balance/test/mocks/interfaces.mock.go

.PHONY: wire
wire:
	@go tool -modfile go.tool.mod wire internal/wire/wire.go

.PHONY: migration # to create new migration
migration:
	@if [ -z "$(name)" ]; then echo "name is required"; exit 1; fi
	touch "internal/db/migrations/$(shell date "+%Y%m%d%H%M%S")_$(name).sql"

.PHONY: prepare
prepare:
	@if [ ! -f go.mod ]; then go mod init $(MODULE); go mod tidy; fi
	@if [ ! -f go.tool.mod ]; then go mod init -modfile go.tool.mod $(MODULE); go mod tidy; fi
	go get -tool -modfile go.tool.mod github.com/air-verse/air@latest
	go get -tool -modfile go.tool.mod github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go get -tool -modfile go.tool.mod github.com/google/wire/cmd/wire@latest
	go get -tool -modfile go.tool.mod go.uber.org/mock/mockgen@latest
	go get -tool -modfile go.tool.mod mvdan.cc/gofumpt@latest

.PHONY: structure
structure:
	@if [ -z "$(db)" ]; then echo "db is required"; exit 1; fi
	@sqlite3 $(db) ".schema --indent" > internal/db/structure.sql
