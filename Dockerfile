FROM golang:1.12.6 as builder
ADD . /go/src/github.com/NV4RE/boilerplate-go/
WORKDIR /go/src/github.com/NV4RE/boilerplate-go
RUN go build

FROM alpine
COPY --from=builder /go/src/github.com/NV4RE/boilerplate-go/dist/boilerplate-go /boilerplate-go
ENTRYPOINT ["./boilerplate-go"]
