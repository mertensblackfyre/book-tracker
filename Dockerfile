# Build stage
FROM golang:1.21-alpine AS builder 
WORKDIR /app
COPY . .  
RUN apk add --no-cache gcc musl-dev
RUN go build -o /app/main

# Run stage 
FROM alpine:latest  
WORKDIR /app
COPY ./.env /app
COPY --from=builder /app/main .
COPY --from=builder /app/static /static
COPY --from=builder /app/.env .

CMD ["./main"]

