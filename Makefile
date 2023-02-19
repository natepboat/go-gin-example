format:
	gofmt -w .
test:
	go test ./...
coverage:
	go test ./... -cover
clean:
	go clean
	rm -r build
compile:
	mkdir -p build
	go build -o build
build: clean test compile
run:
	./build/gin-example
run-local:
	./build/gin-example --env=local