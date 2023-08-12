# Use Go as base image
FROM golang:1.19-alpine
RUN mkdir app
ADD . /app
WORKDIR /app/cmd

RUN go build -o main .
# Command to run app
CMD ["/app/main"]