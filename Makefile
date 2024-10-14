GO=/opt/homebrew/bin/go
LINTER=github.com/golangci/golangci-lint/cmd/golangci-lint@latest
FORMATTER=mvdan.cc/gofumpt@latest
MAIN=cmd/balance/main.go
OUTPUT=balance
RM=rm -f
TEMPL=github.com/a-h/templ/cmd/templ@latest
MOCKGEN=go.uber.org/mock/mockgen@latest
WIRE=github.com/google/wire/cmd/wire@latest

.PHONY: default
default: gen vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@$(GO) vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@$(GO) fix ./...

.PHONY: fmt
fmt: gofmt templfmt

.PHONY: lint
lint:
	@echo "go lint"
	@$(GO) run $(LINTER) run

.PHONY: test
test:
	@echo "go test"
	@$(GO) test ./test/...

.PHONY: run
run: gen vet
	@echo "go run"
	@$(GO) run $(MAIN)

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
gen: wiregen
	@echo "go gen"
	@$(GO) run $(TEMPL) generate

.PHONY: gofmt
gofmt:
			@echo "go fmt"
			@$(GO) run $(FORMATTER) -l -w .

.PHONY: templfmt
templfmt:
			@$(GO) run $(TEMPL) fmt .

.PHONY: mockgen
mockgen:
	@$(GO) run $(MOCKGEN) \
	-source internal/repositories/interfaces.go \
	-package mockedrepositories \
	-destination mocks/repositories/interfaces.mock.go

	@$(GO) run $(MOCKGEN) \
	-source internal/services/interfaces.go \
	-package mockedservices \
	-destination mocks/services/interfaces.mock.go

.PHONY: wiregen
wiregen:
	@$(GO) run $(WIRE) gen internal/server/server.go internal/server/wire.go
