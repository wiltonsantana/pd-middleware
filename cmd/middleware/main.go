package main

import (
	"fmt"
	"os"

	"github.com/joaoaneto/pd-middleware/cmd/middleware/app"

	services "github.com/joaoaneto/pd-middleware/cmd/middleware/common-services"
)

func main() {
	namingService := services.NewNamingService()
	proxy, err := namingService.Lookup("Calculator")
	if err != nil {
		fmt.Printf("naming record not found\n")
		os.Exit(1)
	}

	calculatorProxy := proxy.(app.CalculatorProxy)
	calculatorProxy.Add(1, 4)
}
