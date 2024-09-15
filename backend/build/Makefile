GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=tmp/main
MAIN_PATH=../cmd/api/main.go
MIGRATE_PATH=../cmd/migrate/main.go

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
	./$(BINARY_NAME)

deps:
	$(GOMOD) tidy

lint:
	golangci-lint run

migrate:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MIGRATE_PATH)
	./$(BINARY_NAME)

.PHONY: all build test clean run deps lint