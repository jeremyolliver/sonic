.PHONY: test
test:
	@go test

.PHONY: clean
clean:
	@go fmt && go mod tidy

.PHONY: build
build:
	@go fmt && go mod tidy && go build

.PHONY: install
install:
	@go install
