package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Turtel216/collecix-web-service/cmd/web"
)

func main() {
	app := web.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
