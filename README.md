# wwdatabase-go-backend

## Compiler code

First time you compile the code Update your PATH so that the protoc compiler can find the plugins:
export PATH="$PATH:$(go env GOPATH)/bin"

Then everytime you update the proto file run the following code.

protoc --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 pkg/wwdatabase/wwdatabase.proto

timestamp in go timestamppb.Timestamp
