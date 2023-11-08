FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN make build

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/userapp ./bin/userapp
COPY ./config/prod.yaml ./config/prod.yaml
COPY ./migrations ./migrations
EXPOSE 8000
CMD ["./bin/userapp", "-config", "config/prod.yaml"]
