# ldk-proxy-go

## Proxy

Start the websocket tcp proxy.

```
go run proxy/proxy.go
```

Start the TCP server somewhere (lightning node ideally, or example tcp server `go run main.go`). Use port 19735 or change hardcoded port number.

Connect to the websocket from the browser via `127.0.0.1:8080`

Start passing messages in to the websocket, they will be forwarded to TCP and vis versa.

The example TCP server will respond to messages with a time stamp, useful for testing the proxy and response.
