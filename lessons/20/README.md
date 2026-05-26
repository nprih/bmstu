sudo apt update && sudo apt install -y protobuf-compiler
protoc --version

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest