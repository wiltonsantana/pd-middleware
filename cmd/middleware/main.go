package main

import (
	"github.com/joaoaneto/pd-middleware/cmd/middleware/distribution"
	"github.com/joaoaneto/pd-middleware/cmd/middleware/infrastructure"
)

func main() {
	crh := infrastructure.NewClientRequestHandler("localhost", 4000)
	marshaller := distribution.NewMarshaller()
	requestor := distribution.NewRequestor(marshaller, crh)

	requestor.Invoke()
}
