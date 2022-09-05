## Don't forget to install the following:
## go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
## go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
## Plus you have to rename protoc-gen-go-grpc.exe to protoc-gen-go_grpc.exe on windows machine

protoc -I ../tools/proto --go_out=../../ --go_grpc_out=../../ ../tools/proto/likes.proto