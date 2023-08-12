# Use golang:1.19-alpine as the builder stage
FROM golang:1.19-alpine as builder

# Set /app as the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Copy the .env file from the cmd folder to the current directory
COPY cmd/.env .

# Download dependencies using go mod
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd

# Use alpine:latest as the final stage
FROM alpine:latest

# Set /app as the working directory
WORKDIR /app

# Copy the main executable file from the builder stage
COPY --from=builder /app/main .


# Expose port 8080 to the outside world
EXPOSE 8080

# Run the main executable file
CMD ["./main"]
