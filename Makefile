.PHONY: build install test clean fmt lint run

BINARY=capture
GO=go

build:
	$(GO) build -o $(BINARY) .

install:
	$(GO) install .

test:
	$(GO) test ./... -v

test-coverage:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -html=coverage.out -o coverage.html

clean:
	rm -f $(BINARY) coverage.out coverage.html

fmt:
	gofmt -w .
	goimports -w .

lint:
	golangci-lint run

run: build
	./$(BINARY)
