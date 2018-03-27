FROM golang:1.10-alpine AS builder
# RUN apk add --no-cache build-base go git
WORKDIR /go/src/github.com/chakrit/arcade
ENV GOPATH /go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install -a -ldflags '-extldflags "-static"' ./cmd/arcade

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /root
COPY --from=builder /go/bin/arcade .
ENTRYPOINT ["/root/arcade"]
