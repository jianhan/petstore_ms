# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=store
BINARY_UNIX=$(BINARY_NAME)_unix
DIR_NAME=store

all: test build pb lint

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	./$(BINARY_NAME) --registry=consul

pb:
	for f in proto/**/*.proto; do \
		protoc -I. -I${GOPATH}/src --micro_out=. --go_out=plugins=micro,grpc:. $$f; \
		echo compiled: $$f; \
	done
injectpb:
	for f in proto/**/*.pb.go; do \
		protoc-go-inject-tag -input=$$f \
		echo injecting: $$f; \
	done
lint:
	gometalinter.v2 ./... --vendor