package app

import (
	"github.com/joaoaneto/pd-middleware/cmd/middleware/distribution"
)

// Calculator is an interface to implement calculator proxies.
type Calculator interface {
	Add(param1 int, param2 int) int
}

// CalculatorProxy represents the proxy to Calculator remote object.
type CalculatorProxy struct {
	client distribution.ClientProxy
}

// NewCalculatorProxy creates a brand Calculator instance.
func NewCalculatorProxy(client distribution.ClientProxy) Calculator {
	return CalculatorProxy{client}
}

// Add is an implementation of add operation of Calculator remote object.
func (proxy CalculatorProxy) Add(a int, b int) int {
	requestor := distribution.NewRequestor()
	invocation := distribution.NewInvocation(
		proxy.client.ObjectID,
		proxy.client.Hostname,
		proxy.client.Port,
		"add",
		[]int{a, b},
	)

	requestor.Invoke(invocation)

	return 2 + 2
}
