build/proto:
	buf generate
build/server:
	go build ./server
build/cli:
	go build ./cli
