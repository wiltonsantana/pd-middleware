package infrastructure

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// ClientRequestHandler represents the client request handler component of the middleware.
type ClientRequestHandler struct {
	Hostname   string
	Port       int
	connection net.Conn
}

// NewClientRequestHandler creates a brand ClientRequestHandler instance.
func NewClientRequestHandler(hostname string, port int) ClientRequestHandler {
	crh := ClientRequestHandler{Hostname: hostname, Port: port}
	crh.Connect()
	return crh
}

// Connect establish a connection with the TCP server.
func (crh *ClientRequestHandler) Connect() {
	connection, err := net.Dial("tcp", buildURI(crh.Hostname, crh.Port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	crh.connection = connection
}

func buildURI(hostname string, port int) string {
	return hostname + ":" + strconv.Itoa(port)
}

// Send is a method responsible to send bytes to the TCP server.
func (crh *ClientRequestHandler) Send(data []byte) {
	crh.connection.Write(data)
}
