build/proto:
	buf generate
build/server:
	go build ./server
build/cli:
	go build ./cli

test/server:
	go test -v ./server/tests/.../
