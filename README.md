# ldk-proxy-go

## Proxy

Start the websocket tcp proxy.

```
go run proxy/proxy.go
```

Start the TCP server somewhere (LDK node ideally, or example tcp server `go run main.go`). Use port 19735 or change hardcoded port number.

Connect to the websocket via `127.0.0.1:19735`

Start passing messages in to the websocket, they will be forwarded to TCP and vis versa. If using LDK PeerManager, these should take place automatically when adding a peer to the manager.

The example TCP server will respond to messages with a time stamp.
