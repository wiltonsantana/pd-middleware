package infrastructure

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// ServerRequestHandler represents the server request handler component of the middleware.
type ServerRequestHandler struct {
	Port       int
	connection net.Conn
}

// NewServerRequestHandler creates a brand ServerRequestHandler instance.
func NewServerRequestHandler(port int) ServerRequestHandler {
	srh := ServerRequestHandler{Port: port}
	srh.Connect()
	return srh
}

// Connect establish a connection with the TCP client.
func (srh *ServerRequestHandler) Connect() {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(srh.Port))
	connection, err := listen.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	connection.Close()
	srh.connection = connection
}

// Receive is responsable to receive bytes from the TCP client.
func (srh *ServerRequestHandler) Receive() []byte {
	message := make([]byte, 4096)
	srh.connection.Read(message)
	return message
}

// Send is responsable to send bytes to the TCP client.
func (srh *ServerRequestHandler) Send(data []byte) {
	srh.connection.Write(data)
}
