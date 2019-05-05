package distribution

// ClientProxy represents the proxy to a remote object.
type ClientProxy struct {
	Hostname string
	Port     int
	ObjectID int
}

// NewClientProxy creates a brand ClientProxy instance.
func NewClientProxy(hostname string, port int, objectID int) ClientProxy {
	return ClientProxy{Hostname: hostname, Port: port, ObjectID: objectID}
}
