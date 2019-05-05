package main

import "github.com/joaoaneto/pd-middleware/cmd/middleware/app"

func main() {
	calculatorProxy := app.NewCalculatorProxy()
	calculatorProxy.Add(1, 4)
}
