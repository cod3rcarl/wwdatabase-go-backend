
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

run:
	DB_CONNECTION="user=mnjghfpgupfubo password=2fcdf645897feac50ddb34cd0ae211e064636c95f8453899ed289eeb183ae70a host=ec2-54-155-112-143.eu-west-1.compute.amazonaws.com port=5432 dbname=d3leb3ukf4fo82 pool_max_conns=10" \
	GRPC_HOST=127.0.0.1 \
	GRPC_PORT=3000 \
	LOGGER_LEVEL=debug \
	go run cmd/app/main.go
