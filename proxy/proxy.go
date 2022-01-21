package main

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func address_decode(address_bin []byte) (string, string) {
	var host string = "127.0.0.1"
	var port string = "19735"

	return host, port
}

func forwardtcp(wsconn *websocket.Conn, conn net.Conn) {

	for {
		// Receive and forward pending data from tcp socket to web socket
		tcpbuffer := make([]byte, 1024)

		n, err := conn.Read(tcpbuffer)
		if err == io.EOF {
			fmt.Printf("TCP Read failed")
			break
		}
		if err == nil {
			fmt.Printf("Forwarding from tcp to ws: %d bytes: %s\n", n, tcpbuffer)
			// print_binary(tcpbuffer)
			wsconn.WriteMessage(websocket.BinaryMessage, tcpbuffer[:n])
		}
	}
}

func forwardws(wsconn *websocket.Conn, conn net.Conn) {

	for {
		// Send pending data to tcp socket
		n, buffer, err := wsconn.ReadMessage()
		if err == io.EOF {
			fmt.Printf("WS Read Failed")
			break
		}
		if err != nil {
			panic(err)
		}

		s := string(buffer[:len(buffer)])
		fmt.Printf("Received (from ws) forwarding to tcp: %d bytes: %s %d\n", len(buffer), s, n)
		fmt.Fprintf(conn, s+"\n")
		if err != nil {
			panic(err)
		}
	}
}

func wsProxyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received websocket connection")

	wsconn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		panic(err)
	}

	// get connection address and port
	address := make([]byte, 16)

	/*
		n, address, err := wsconn.ReadMessage()
		if err != nil {
			fmt.Printf("address read error")
			fmt.Printf("read %d bytes", n)
		}

	*/

	host, port := address_decode(address)

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Started tcp connection to " + host + ":" + port)

	go forwardtcp(wsconn, conn)
	go forwardws(wsconn, conn)

	fmt.Println("forwarding started")
}

func main() {
	fmt.Println("Starting websocket: 127.0.0.1:8080/proxy")

	http.HandleFunc("/proxy", wsProxyHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
