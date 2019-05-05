package distribution

// Invocation represents the object remote request data.
type Invocation struct {
	objectID      int
	ipAddress     string
	port          int
	operationName string
	parameters    []int
}

// NewInvocation creates a brand Invocation instance.
func NewInvocation(objectID int, ipAddress string, port int, operationName string, parameters []int) Invocation {
	return Invocation{objectID, ipAddress, port, operationName, parameters}
}
