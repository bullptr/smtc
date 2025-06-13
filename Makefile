.PHONY: all build test bench

all: build

build:
	@clear
	@go build

test-001:
	@clear
	@go run . check ./tests/test-001.smt2

test:
	@clear
	@go test -v ./...

bench:
	@clear
	@go test -v ./benchmarks/... -bench=.
