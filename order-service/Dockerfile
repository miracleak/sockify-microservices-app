FROM golang:1.20

WORKDIR /order-service

COPY /main .

RUN go mod init github.com/miracleak/sockify-microservices-app

RUN go mod tidy 

RUN go build main.go

EXPOSE 8080

CMD ["./main"]