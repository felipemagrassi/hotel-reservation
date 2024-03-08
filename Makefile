.PHONY: build
build:
	@go build -o bin/api

.PHONY: run
run: build
	@./bin/api

.PHONY: test
test:
	@go test -v ./...


