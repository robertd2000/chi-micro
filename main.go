package main

import (
	"context"
	"fmt"

	"github.com/robertd2000/orders-api/app"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Errorf("Failed with %w", err)
	}
}
