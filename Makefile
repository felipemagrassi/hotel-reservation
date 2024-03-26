.PHONY: build
build:
	@go build -o bin/api

.PHONY: run
run: build
	@./bin/api

.PHONY: seed
seed: 
	@go run scripts/seed.go

.PHONY: test
test:
	@go test -v ./...


