package distribution

import "encoding/json"

// Marshaller provides mapping and serializing capabilities to support distribution layer.
type Marshaller struct {
}

// NewMarshaller creates a brand marshaller instance.
func NewMarshaller() Marshaller {
	return Marshaller{}
}

// Marshal maps the message to JSON and serialize it in bytes.
func (marshaller *Marshaller) Marshal(message Message) ([]byte, error) {
	return json.Marshal(message)
}

// Unmarshal deserialize bytes message and maps it to JSON.
func (marshaller *Marshaller) Unmarshal(message []byte) (*Message, error) {
	msg := Message{}
	err := json.Unmarshal(message, &msg)
	return &msg, err
}
