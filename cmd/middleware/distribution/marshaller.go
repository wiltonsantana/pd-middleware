package distribution

import "encoding/json"

// Marshaller provides mapping and serializing capabilities to support distribution layer.
type Marshaller struct {
}

// NewMarshaller creates a brand marshaller instance.
func NewMarshaller() Marshaller {
	return Marshaller{}
}

// Marshal maps the message to JSON e serialize it in bytes.
func (marshaller *Marshaller) Marshal(message Message) ([]byte, error) {
	return json.Marshal(message)
}
