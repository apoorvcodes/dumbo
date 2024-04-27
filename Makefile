
all: say_version build test clean

tests: build test clean

say_version:
	@echo "apoorvcodes/dumbo v1"

build:
	@echo "Building..."
	go build dumbo.go

test:
	@echo "Testing..."
	go test

clean:
	@echo "Cleaning up..."
	go fmt ./