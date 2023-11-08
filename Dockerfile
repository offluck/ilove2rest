FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o userapp ./...

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/userapp .
EXPOSE 8000
CMD ["./userapp"]
