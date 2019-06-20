# api service

FROM golang:alpine as builder

LABEL maintainer="Jian Han <jianhan@mail.com>"

RUN apk --no-cache add git

WORKDIR /store

COPY . .

RUN go get -d -v ./...

RUN go build -o store

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /store/store .

# Run the binary as per usual! This time with a binary build in a
# separate container, with all of the correct dependencies and
# run time libraries.
ENTRYPOINT [ "./store" ]