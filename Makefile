build-wasm:
	GOOS=js GOARCH=wasm go build -o showcase/web/app.wasm ./showcase/

run: build-wasm
	cd showcase && go run .

# air --build.pre_cmd="GOOS=js GOARCH=wasm go build -o ./web/app.wasm"