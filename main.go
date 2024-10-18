package main

import (
	"GoPractice/HTTP_SERVER/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8080")

	go srv.ListenAndServe()
	log.Println("Server Started...")
	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("Server Stoped...")
}
