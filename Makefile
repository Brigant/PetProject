# Declare variables
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOLINTCMD=golangci-lint
LINTRUN=$(GOLINTCMD) run
ENTRYPOINT=backend/cmd/main.go
BINARY_NAME=petproject

# Define targets

all: tests build

run:
	$(GORUN) $(ENTRYPOINT)

all-tests:
	$(LINTRUN) --enable-all --no-config

tests:
	$(LINTRUN)

build:
	$(GOBUILD) -o $(BINARY_NAME) $(ENTRYPOINT)

clean:
	rm -f $(BINARY_NAME)
