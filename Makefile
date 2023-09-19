build-wasm:
	GOOS=js GOARCH=wasm go build -ldflags="-w -s" -o ./public/main.wasm ./cmd/wasm

run-server:
	go run cmd/server/main.go

