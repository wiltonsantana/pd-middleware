package distribution

import (
	"github.com/joaoaneto/pd-middleware/cmd/middleware/infrastructure"
)

// Requestor represents the "maestro" of distribution layer.
type Requestor struct {
	marshaller Marshaller
	crh        infrastructure.ClientRequestHandler
}

// NewRequestor creates a brand requestor instance.
func NewRequestor() Requestor {
	return Requestor{}
}

// Invoke create a encapsulated message and does the request to server.
func (requestor *Requestor) Invoke(invocation Invocation) {
	requestor.marshaller = NewMarshaller()
	requestor.crh = infrastructure.NewClientRequestHandler(invocation.ipAddress, invocation.port)
	message := createMessage(invocation.objectID, invocation.operationName, invocation.parameters)
	messageMarshalled, _ := requestor.marshaller.Marshal(message)
	requestor.crh.Send(messageMarshalled)
}

func createMessage(objectID int, operation string, parameters []int) Message {
	requestHeader := RequestHeader{0, "", true, objectID, operation}
	requestBody := RequestBody{parameters}

	messageHeader := MessageHeader{"MIOP", 0, false, 0, 0}
	messageBody := MessageBody{requestHeader: requestHeader, requestBody: requestBody}

	return Message{messageHeader, messageBody}
}
