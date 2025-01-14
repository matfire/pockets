build/proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shared/pockets.proto
build/server:
	go build ./server
build/cli:
	go build ./cli
