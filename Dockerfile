FROM golang:1.25.3-alpine3.22 AS builder

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" cmd/main.go



FROM alpine:3.22

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder --chown=appuser:appgroup --chmod=755 /app/main /main
COPY --from=builder --chown=appuser:appgroup --chmod=755 /app/static /static

USER appuser

EXPOSE 8080

CMD ["/main"]
