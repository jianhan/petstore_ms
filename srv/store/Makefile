# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=store
BINARY_UNIX=$(BINARY_NAME)_unix
DIR_NAME=store

all: mockgen test build pb lint

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

mockgen:
	mockgen -destination=./mock/petdatastore_mock.go -package=mock github.com/jianhan/petstore_ms/srv/store/datastore PetDataStore
	mockgen -destination=./mock/petservice_mock.go -package=mock github.com/jianhan/petstore_ms/srv/store/proto/pet PetService
	mockgen -destination=./mock/petservicehandler_mock.go -package=mock github.com/jianhan/petstore_ms/srv/store/proto/pet PetServiceHandler

lint:
	gometalinter.v2 ./... --vendor