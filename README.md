# Verve Assignment

## Overview

This project provides a high-performance REST service to process unique requests with the following features:

## Project Structure
```bash
    ├── ASSIGNMENT.MD
    ├── Dockerfile
    ├── LICENSE
    ├── README.md
    ├── controllers
    │   ├── acceptController.go
    │   └── uniqueRequestsController.go
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── jobs
    │   └── jobScheduler.go
    ├── main.go
    ├── models
    │   └── redisModel.go
    ├── nginx.conf
    ├── routes
    │   └── router.go
    └── utils
        ├── logger.go
        ├── rabbitMQClient.go
        └── redisClient.go
```
## Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/subarna-sahoo/verve_assignment.git
    cd verve_assignment
    ```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```

3. **Run the application:**

    To run the application locally:

    ```bash
    go run main.go
    ```

    This will start the server on `http://localhost:5000`

4. **Dockerize the application:**

    To run the application with Docker:

    ```bash
    docker-compose up --build
    ```

    This will spin up the Redis, RabbitMQ, application and Nginx in containers.

## Endpoints

- `GET /`: Health check.
- `GET /api/verve/accept?id=<id>&endpoint=<endpoint>`: Accepts the request with the provided ID and optional `endpoint` as query, processes the request, and ensures ID uniqueness.
- `GET /api/verve/unique-requests`: Returns the count of unique requests received in the last minute.

## Redis

Ensure Redis is running locally or configure the connection in the `models/redis.go` file. The application uses Redis to store the request IDs and ensure uniqueness.

## RabbitMQ Integration

RabbitMQ is used to send the unique request count for the last minute. This allows for real-time processing of the request count, such as logging or forwarding to other services.

### Steps to Set Up RabbitMQ:
1. RabbitMQ is configured in the `docker-compose.yml` file. It runs as a service and is accessible at the default AMQP port `5672`.
2. The `PublishToRabbitMQ` function sends the count of unique requests to the `unique_requests` queue every minute.
3. We can setup a worker to listens to this queue and processes the messages.


## Nginx Load Balancing

Nginx is used to load balance between multiple instances of the backend application. It ensures that the incoming requests are evenly distributed across the available backend instances.

The Nginx configuration uses an `upstream` directive to define two backend servers (e.g., `go_app1` and `go_app2`). These instances will process requests in parallel and share the load.
