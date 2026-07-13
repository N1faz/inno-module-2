FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app/main \
    main.go


FROM gcr.io/distroless/static

COPY --from=builder /app/main app

CMD ["/app"]
