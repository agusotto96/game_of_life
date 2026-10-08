GOROOT := $(shell go env GOROOT)

build-web:
	mkdir -p public
	cp -f web/index.html public/
	cp -f "$(GOROOT)/lib/wasm/wasm_exec.js" public/
	GOOS=js GOARCH=wasm go build -o public/game.wasm .

serve: build-web
	go run ./cmd/serve
