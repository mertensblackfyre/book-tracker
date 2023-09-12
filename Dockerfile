# Run stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

# Build stage
FROM alpine:latest
WORKDIR /app
COPY ./.env /app
COPY --from=builder /app/main .



EXPOSE 8080
CMD ["/app/main"]
