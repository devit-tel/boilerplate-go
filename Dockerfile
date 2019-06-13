FROM golang:1.12.6 as builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ADD . /go/src/github.com/NV4RE/boilerplate-go/
WORKDIR /go/src/github.com/NV4RE/boilerplate-go
RUN make build

FROM alpine
COPY --from=builder /go/src/github.com/NV4RE/boilerplate-go/dist/boilerplate-go /boilerplate-go
ENTRYPOINT ["./boilerplate-go"]
