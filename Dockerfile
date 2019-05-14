FROM golang:1.12-alpine3.9 AS build
RUN apk add --update git make gcc musl-dev curl ca-certificates openssl git openssh file
ENV GOPATH /go
ENV PATH /go/bin:$PATH
# Add source code
ADD . /go/src/github.com/acottais/portctl/
WORKDIR /go/src/github.com/acottais/portctl/
RUN mkdir -p ./build && make build


FROM alpine:latest
RUN apk add --no-cache --update ca-certificates openssl
# server-side dirs
RUN addgroup -S portctl && \
    adduser -S portctl -G portctl
COPY --from=build /go/src/github.com/acottais/portctl/build/* /usr/local/bin/
