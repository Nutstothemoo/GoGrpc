$ export PATH="$PATH:$(go env GOPATH)/bin" // Permet de trouver le plugin compiler go
$ protoc --go_out=. --go-grpc_out=. proto/greet.proto // compile en go

