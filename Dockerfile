# Build
FROM golang:1.19-alpine3.15 as builder


WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build app ./cmd

# Run
FROM alpine:3.15

WORKDIR /root/
COPY --from=builder /src/app .
CMD ["./app"]
