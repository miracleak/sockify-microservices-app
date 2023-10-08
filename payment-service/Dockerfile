# Builder stage
FROM golang:1.20 as builder

WORKDIR /payment-service

COPY main.go .

RUN go mod init github.com/miracleak/sockify-microservices-app

RUN go mod tidy 

RUN go build -o main

# Final stage
FROM alpine:latest

ENV ORDER_SERVICE_HOST order-service
ENV ORDER_SERVICE_PORT 8080

WORKDIR /app

# Copy the compiled Go application from the builder stage to the final image
COPY --from=builder /payment-service/main /app/main

EXPOSE 8081

CMD ["./main"]
