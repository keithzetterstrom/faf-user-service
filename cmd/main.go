package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/keithzetterstrom/faf-user-service/cmd/app"
)

func main() {
	b := app.Start()
	defer b.Close()

	go b.ServerStart()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
}
