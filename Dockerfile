FROM golang:1.13.0-alpine3.10 as builder
COPY *.go /app/
WORKDIR /app
ARG CGO_ENABLED=0
RUN go build -o manitool manitool.go

FROM scratch
COPY --from=golang:1.13 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=golang:1.13 /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/manitool /app/
WORKDIR /app
