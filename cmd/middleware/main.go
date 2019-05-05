package main

import (
	"github.com/joaoaneto/pd-middleware/cmd/middleware/infrastructure"
)

func main() {
	crh := infrastructure.ClientRequestHandler{Hostname: "localhost", Port: 4000}
	crh.Connect()
	crh.Send([]byte("Hello World"))
}
