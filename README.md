# Go WebSocket Server

A minimal WebSocket echo server written in Go, packaged with Docker.
Includes a simple HTML test page served at `/` so you can connect and test directly in your browser.

---

## Features

- WebSocket endpoint at **`/ws`**
- Echoes messages back to the client
- JSON-structured logs with timestamps, log level, and client IP
- Built-in HTML test page at **`/`**

---

## Run locally (without Docker)

```bash
go run main.go
```

Open [http://localhost:8080](http://localhost:8080) in your browser.

---

## Build & Run with Docker

```bash
# build image
docker build -t go-ws-server .

# run container
docker run -p 8080:8080 go-ws-server
```

---

## Test WebSocket

### Browser
Go to: [http://localhost:8080](http://localhost:8080)
Use the test page to connect, send, and receive messages.

### CLI
Using [websocat](https://github.com/vi/websocat):

```bash
websocat ws://localhost:8080/ws
```

Type any message — it will be echoed back.

---

## Logs

Logs are JSON-formatted:

```json
{"level":"info","time":"2025-09-18T08:55:10Z","message":"Client connected","client":"172.17.0.1"}
{"level":"info","time":"2025-09-18T08:55:15Z","message":"Received: hello","client":"172.17.0.1"}
```

---
