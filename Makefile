build/proto:
	npx buf generate
build/server:
	go build ./server
build/cli:
	go build ./cli

test/server:
	go test -v ./server/tests/.../

dev/server:
	go run -tags dev server/main.go

dev/web:
	cd server/web && pnpm dev