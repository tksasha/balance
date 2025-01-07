GO=/opt/homebrew/bin/go
MAIN=cmd/balance/main.go
OUTPUT=balance
RM=rm -f

AIR=github.com/air-verse/air@latest
FORMATTER=mvdan.cc/gofumpt@latest
LINTER=github.com/golangci/golangci-lint/cmd/golangci-lint@latest
MOCKGEN=go.uber.org/mock/mockgen@latest
WIRE=github.com/google/wire/cmd/wire@latest

.PHONY: default
default: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@$(GO) vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@$(GO) fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@$(GO) run $(FORMATTER) -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@$(GO) run $(LINTER) run

.PHONY: test
test:
	@echo "go test"
	@$(GO) test ./test/...

# .PHONY: run
# run:
# 	@echo "go run"
# 	@$(GO) run $(AIR)

.PHONY: run
run:
	@echo "go run (without live reloading)"
	@$(GO) run github.com/tksasha/balance/cmd/balance

.PHONY: build
build:
	@echo "go build"
	@$(GO) build -o $(OUTPUT) $(MAIN)

.PHONY: clear
clear:
	@echo "go clear"
	@$(RM) $(OUTPUT)

.PHONY: clean
clean: clear

.PHONY: gen
gen: wire mockgen

.PHONY: mockgen
mockgen:
	@$(GO) run $(MOCKGEN) \
		-source internal/services/interfaces.go \
		-package mocksforservices \
		-destination mocks/services/interfaces.mock.go

	@$(GO) run $(MOCKGEN) \
		-source internal/handlers/interfaces.go \
		-package mocksforhandlers \
		-destination mocks/handlers/interfaces.mock.go

.PHONY: wire
wire:
	@echo "wire gen"
	@$(GO) run $(WIRE) internal/wire/wire.go
