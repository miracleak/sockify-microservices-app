# Sockify Microservices App

Welcome to the Sockify Microservices App! This is a multi-service application that simulates a full-featured e-commerce platform for socks. Each microservice handles a specific aspect of the application, making it scalable, maintainable, and easy to manage.

![Sockify Microservices App](./images/socks-app.png)

## Table of Contents

- [Overview](#overview)
- [Microservices](#microservices)
- [Getting Started](#getting-started)
- [Architecture](#architecture)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [License](#license)

## Overview

The Sockify Microservices App is built using Go and follows best practices for microservices architecture. It consists of the following microservices:

1. **User Service**: Manages user registration, authentication, and profiles.
2. **Product Service**: Handles sock catalog, pricing, and product details.
3. **Order Service**: Manages order creation, payment processing, and order status.
4. **Payment Service**: Processes payments for orders.
5. **Shipping Service**: Handles shipping and delivery logistics.

## Microservices

Each microservice is contained in its respective directory:

- `/user-service`
- `/product-service`
- `/order-service`
- `/payment-service`
- `/shipping-service`

Navigate to each microservice's directory to find more information about its structure and functionality.

## Getting Started

To run the Socks Microservices App on your local machine, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/socks-microservices-app.git

## Navigate to the root directory:
    ```bash
    cd socks-microservices-app


## Build and start each microservice:
    ```bash 
    cd user-service && go run main.go
    cd ../product-service && go run main.go
    cd ../order-service && go run main.go
    cd ../payment-service && go run main.go
    cd ../shipping-service && go run main.go

 Access the application at http://localhost:8080.

   Architecture
The Socks Microservices App follows a microservices architecture, which provides scalability, flexibility, and ease of maintenance. Each microservice is designed to be independent and communicates with others through RESTful APIs.


API Documentation
For detailed API documentation and usage instructions, refer to the API documentation for each microservice. You can find the documentation in the respective microservice directories.

 Deployment
To deploy this application in a production environment, refer to the deployment guides in each microservice's directory. You may need to set up a container orchestration platform like Kubernetes or Docker Compose

