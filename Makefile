.PHONY: test build clean run

test:
	go test ./...

build: 
	@mkdir -p build/casem
	go build -o build/casem/casem ./cmd/casem.go

clean: 
	rm -rf build

run: test build
	./build/casem/casem
