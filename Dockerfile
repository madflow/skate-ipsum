FROM golang:1.20-alpine AS builder

WORKDIR /build

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" . 

FROM scratch

COPY --from=builder /build/skate-ipsum /usr/local/bin/skate-ipsum

ENTRYPOINT ["/usr/local/bin/skate-ipsum"]

EXPOSE 6969


