FROM golang:1.12.6 as builder
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GIN_MODE=release
ADD . /go/src/github.com/devit-tel/boilerplate-go/
WORKDIR /go/src/github.com/devit-tel/boilerplate-go
RUN make build

FROM alpine
ENV GIN_MODE=release
COPY --from=builder /go/src/github.com/devit-tel/boilerplate-go/dist/boilerplate-go ./
ENTRYPOINT ["/bin/sh", "-c", "sleep 99999999"]
