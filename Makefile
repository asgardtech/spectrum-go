build-wasm:
	GOOS=js GOARCH=wasm go build -o showcase/web/app.wasm ./showcase/

run: build-wasm
	cd showcase && go run .

watch:
	air
