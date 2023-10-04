
Installation :
- brew install protoc
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"


Converting Proto File to Go File
- protoc --go_out=. --go-grpc_out=. proto/calculator.proto
