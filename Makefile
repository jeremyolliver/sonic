.PHONY: test
test:
	@go test

.PHONY: build
build:
	@go fmt && go mod tidy && go build -o sonic
