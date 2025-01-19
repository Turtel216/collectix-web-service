package main

import (
	"context"
	"fmt"

	"github.com/Turtel216/Go-Microservice/cmd/web"
)

func main() {
	app := web.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

	fmt.Println("Started Application at 8080")
}
