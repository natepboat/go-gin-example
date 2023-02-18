format:
	gofmt -w .
test:
	go test ./...
coverage:
	go test ./... -cover
clean:
	rm -r build
compile:
	mkdir -p build
	go build -o build
build: clean test compile
run:
	./build/gin-example