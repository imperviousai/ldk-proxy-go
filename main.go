package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	// Startup the HTTP server for receiving proxy messages
	h := HttpHandler{}

	httpServer := &http.Server{
		Addr:    "127.0.0.1:8889",
		Handler: h,
	}

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	// TODO wg
	h.Listen("19735")
}

// create an http handler for the proxy server
type HttpHandler struct{}

func (h HttpHandler) Listen(portArg string) {
	// Listen on TCP port
	PORT := ":" + portArg
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Accepted a connection")

	for {
		fmt.Println("Waiting for a response...")

		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}

}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")
	res.Write(data)
}
