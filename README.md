# Verve Assignment

## Overview

This project provides a high-performance REST service to process unique requests with the following features:
- `/api/verve/accept`: Accepts an ID and an optional endpoint, processes the request, and ensures the uniqueness of the ID using Redis.
- Logs the number of unique requests processed every minute.
- Optionally sends the unique request count to a specified endpoint.

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

    ```bash
    go run main.go
    ```

    This will start the server on `http://localhost:8080`.

4. **Dockerize the application:**

    ```bash
    docker-compose up --build
    ```

## Endpoints

- `GET /api/ping`: Health check.
- `GET /api/verve/accept?id=<id>&endpoint=<endpoint>`: Process a request.
- `GET /api/verve/stats`: Get unique request count.

## Redis

Make sure you have Redis running locally or configure the connection in the `models/redis.go` file.
