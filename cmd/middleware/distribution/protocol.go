package distribution

// Message represents the session protocol message.
type Message struct {
	header MessageHeader
	body   MessageBody
}

// MessageHeader represents the session protocol message headers.
type MessageHeader struct {
	magic       string
	version     int
	byteOrder   bool
	messageType int
	messageSize int64
}

// MessageBody represents the session protocol message body.
type MessageBody struct {
	requestHeader RequestHeader
	requestBody   RequestBody
	replyHeader   ReplyHeader
	replyBody     ReplyBody
}

// RequestHeader represents the request headers.
type RequestHeader struct {
	id               int
	context          string
	responseExpected bool
	objectKey        int
	operation        string
}

// RequestBody represents the request body.
type RequestBody struct {
	parameters []int
}

// ReplyHeader represents the reply headers.
type ReplyHeader struct {
	serviceContext string
	requestID      int
	replyStatus    int
}

// ReplyBody represents the reply body.
type ReplyBody struct {
	operationResult int
}
