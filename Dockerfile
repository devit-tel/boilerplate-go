FROM golang:1.12.6 as builder
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GIN_MODE=release
ADD . /go/src/github.com/NV4RE/boilerplate-go/
WORKDIR /go/src/github.com/NV4RE/boilerplate-go
RUN make build

FROM alpine
COPY --from=builder /go/src/github.com/NV4RE/boilerplate-go/dist/boilerplate-go ./
ENTRYPOINT ["./boilerplate-go"]
