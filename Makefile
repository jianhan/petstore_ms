.PHONY: proto build

proto:
	for d in api srv; do \
		for f in $$d/**/proto/*.proto; do \
			protoc -I. -I${GOPATH}/src --micro_out=. --go_out=plugins=micro,grpc:. $$f; \
			echo compiled: $$f; \
		done \
	done

lint:
	./bin/lint.sh

build:
	./bin/build.sh

run:
	docker-compose build
	docker-compose up
