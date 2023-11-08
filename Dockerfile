FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN make build

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/userapp .
COPY ./config/prod.yaml ./prod.yaml
EXPOSE 8000
CMD ["./userapp", "-config", "prod.yaml"]
