package distribution

import (
	"fmt"

	"github.com/joaoaneto/pd-middleware/cmd/middleware/infrastructure"
)

// Requestor represents the "maestro" of distribution layer.
type Requestor struct {
	marshaller Marshaller
	crh        infrastructure.ClientRequestHandler
}

// NewRequestor creates a brand requestor instance.
func NewRequestor(marshaller Marshaller, chr infrastructure.ClientRequestHandler) Requestor {
	return Requestor{marshaller, chr}
}

func (requestor *Requestor) Invoke() {
	message := createMessage()
	fmt.Println(message)
	messageMarshalled, _ := requestor.marshaller.Marshal(message)
	requestor.crh.Send(messageMarshalled)
}

func createMessage() Message {
	requestHeader := RequestHeader{0, "", true, 0, "add"}
	requestBody := RequestBody{[]int{1, 2}}

	messageHeader := MessageHeader{"MIOP", 0, false, 0, 0}
	messageBody := MessageBody{requestHeader: requestHeader, requestBody: requestBody}

	return Message{messageHeader, messageBody}
}
