
ifeq ($(OS),Windows_NT)
EXT = .exe
else
EXT =
endif

OUT = grpc/pkg/wwdatabase

install-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto:
	protoc --go_out=${OUT} --go_opt=paths=source_relative --go-grpc_out=${OUT} --go-grpc_opt=paths=source_relative grpc/pkg/wwdatabase/wwdatabase.proto

generate:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

server: ## Run the main application
	go run grpc/cmd/main.go

client: ## Run the main application
	go run graphql/cmd/app/main.go

tidy: ## Tidy up Go modules
	go mod tidy

cilint: ## Run linters
	golangci-lint$(EXT) run -v

gofumpt: ## Format code with gofumpt
	gofumpt$(EXT) -l -w .

check: gofumpt cilint tidy ## Format code, run linters and tidy up Go modules
