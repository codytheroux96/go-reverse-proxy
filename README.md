# Go Reverse Proxy 

This project is a basic reverse proxy built using the Go standard library. 

## Features

- Listening for incoming HTTP(S) requests
- Routing requests to test backends based on wildcard paths (`/s1/*`, `/s2/*`)
- Forwarding HTTP requests to appropriate backends, preserving method, headers, and body
- Returning backend responses to clients
- Rate limiting to restrict how frequently clients can send requests
- In-memory caching for GET requests
- Error handling with retries if a backend fails
- Timeout handling to prevent hanging requests
- Substantial logging for observability
- HTTPS support with local certificates

## Project Structure

- `main.go`: Starts the proxy and both test servers
- `internal/app/`: Core logic for the proxy (routing, caching, rate limiting)
- `test_servers/server_one/`: A minimal backend responding to `/s1/*` routes
- `test_servers/server_two/`: A second backend for `/s2/*` routes

## Usage

**If you want to run this yourself you will need to either get rid of the HTTPS redirect logic in `main.go` or generate your own certificates for HTTPS support**

Run the main application and the proxy will start on port `:8080`, with test backends on ports `:4200` and `:2200`.

Example routes to test:
- `GET /s1/s1health` - simple GET request with no substance
- `POST /s1/s1list` - simple POST request with no substance
- `GET /s2/s2health` - simple GET request with no substance
- `POST /s2/s2list` - simple POST request with no substance
- `POST /s1/echo` – echoes back the request body
- `POST /s2/echo` – echoes back the request body
- `GET /s1/headers` – returns request headers
- `GET /s2/headers` – returns request headers

For example: `http://localhost:8080/s2/s2headers` will call `http://localhost:2200/s2headers`

## Notes

- This proxy only runs locally; it is **not deployed** and not accessible from outside your machine
- Both the proxy and the backend servers listen on `localhost`
- Routes are currently **hardcoded**, not dynamically configurable