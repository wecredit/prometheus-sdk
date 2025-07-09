# ------------ Multi-Stage Dockerfile for API Build ------------

# Stage 1: Build Stage
FROM golang:1.23.5 AS builder

WORKDIR /go/src/prometheus-sdk
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main ./cmd/main.go

# Stage 2: Run Stage
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=UTC

WORKDIR /app
COPY --from=builder /go/bin/main .
COPY sdk_config.json .

EXPOSE 2112

CMD ["./main"]
