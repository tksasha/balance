GO=/opt/homebrew/bin/go
LINTER=github.com/golangci/golangci-lint/cmd/golangci-lint@latest
FORMATTER=mvdan.cc/gofumpt@latest
MAIN=cmd/balance/main.go
OUTPUT=balance
RM=rm -f
TEMPL=github.com/a-h/templ/cmd/templ@latest

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

.PHONY: run
run:
	@$(GO) run $(MAIN)

.PHONY: build
build:
	@echo "go build"
	@$(GO) build -o $(OUTPUT) $(MAIN)

.PHONY: clear
clear:
	@echo "go clear"
	@$(RM) $(OUTPUT)

.PHONY: gen
gen:
	@$(GO) run $(TEMPL) generate
