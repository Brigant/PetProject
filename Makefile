# Declare variables
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOLINTCMD=golangci-lint
LINTRUN=$(GOLINTCMD) run
TESTRUN=$(COMCD) test
ENTRYPOINT=backend/cmd/main.go
BINARY_NAME=petproject

# Define targets

all: tests build

run:
	$(GORUN) $(ENTRYPOINT)

lintckeck-all:
	$(LINTRUN) --enable-all --no-config

lintcheck:
	$(LINTRUN)

build:
	$(GOBUILD) -o $(BINARY_NAME) $(ENTRYPOINT)

test:
	$(TESTRUN) ./...

clean:
	rm -f $(BINARY_NAME)

coverage:
	go test ./... -coverprofile=coverage.out && $(GOCMD) tool cover -html=coverage.out -o coverage.html	&& xdg-open ./coverage.html
