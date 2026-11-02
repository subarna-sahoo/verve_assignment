# Verve Assignment: Unique Request Counter

A high-throughput REST service in **Go (Gin)** that accepts requests carrying an `id` and counts how many **unique** IDs arrive each minute. It keeps counts correct even when several instances run behind a load balancer. Every minute it publishes the count to **RabbitMQ**.

The original brief is in [ASSIGNMENT.MD](ASSIGNMENT.MD).

## Architecture

```
            ┌──────────┐
 client ──▶ │  Nginx   │  :80  (round-robin)
            └────┬─────┘
          ┌──────┴──────┐
          ▼             ▼
     go_app1:5000   go_app2:5000      Gin service (2 replicas)
          │             │
          └──────┬──────┘
                 ▼
     Redis sorted set "req:ids"       shared dedup across instances
                 │
     every minute: count IDs seen in the last 60 s
                 ▼
     RabbitMQ queue "unique_requests"
```

**Design choices**

- **Deduplication across instances**: each ID is written to a Redis **sorted set** (`ZADD req:ids <unix_ts> <id>`). A set holds each member only once, so two instances receiving the same ID at the same moment still count it once.
- **Per-minute count**: a background ticker runs `ZRANGEBYSCORE` over the last 60 seconds and takes the length of the result.
- **Streaming instead of a log file** (Extension 2): the count is published to RabbitMQ, where any consumer can process it.
- **Non-blocking callbacks**: when `endpoint` is provided, the POST to it runs in a goroutine through a bounded worker pool, so `/accept` returns immediately. The response status code is logged.

## Endpoints

| Method | Path | Description |
|---|---|---|
| `GET` | `/` | Health check. Returns the IP of the instance that served it, which is handy for seeing load balancing at work. |
| `GET` | `/api/verve/accept?id=<int>&endpoint=<url>` | Records `id`. Returns `ok` or `failed`. If `endpoint` is set, sends it a POST. |
| `GET` | `/api/verve/unique-requests` | Debug endpoint: unique IDs seen in the last minute. |

## Running

The service connects to Redis at `redis:6379` and RabbitMQ at `rabbitmq:5672` (hostnames from Docker Compose), so the simplest way to run everything is:

```bash
docker-compose up --build
```

This starts Redis, RabbitMQ (management UI at <http://localhost:15672>, login `guest`/`guest`), two app instances (`:8080` and `:8081`) and Nginx on `:80`.

```bash
curl "http://localhost/api/verve/accept?id=42"
curl "http://localhost/api/verve/unique-requests"
```

To run the Go service on its own (`go run main.go`, port 5000), first change the Redis and RabbitMQ addresses in `utils/redisClient.go` and `utils/rabbitMQClient.go` to `localhost`.

## Project structure

```
main.go                       Starts Redis, RabbitMQ, the minute job and the Gin server
routes/router.go              Route table
controllers/                  /accept, /unique-requests and health handlers
models/redisModel.go          Store an ID, fetch IDs from the last minute
jobs/jobScheduler.go          Ticker that publishes the count every minute
utils/                        Redis and RabbitMQ clients, logger
Dockerfile, docker-compose.yml, nginx.conf
```

## Tech stack

Go 1.24 · Gin · Redis (go-redis v9) · RabbitMQ (amqp) · Nginx · Docker Compose

## License

MIT. See [LICENSE](LICENSE).
