
ifeq ($(OS),Windows_NT)
EXT = .exe
else
EXT =
endif

OUT = grpc/pkg/wwdatabase

tidy: ## Tidy up Go modules
	go mod tidy

cilint: ## Run linters
	golangci-lint$(EXT) run -v

gofumpt: ## Format code with gofumpt
	gofumpt$(EXT) -l -w .

check: gofumpt cilint tidy ## Format code, run linters and tidy up Go modules

# include environment variables needed for DB_CONNECTION | GRPC_HOST | GRPC_PORT | LOGGER_LEVEL
run:
	go run cmd/app/main.go
