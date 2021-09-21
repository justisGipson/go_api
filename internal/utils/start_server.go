package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// StartServer starts the server and shutsdown gracefully
func StartServer(a *fiber.App) {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // catch signals from OS
		<-sigint
		// if interrupt signal received, shutdown
		if err := a.Shutdown(); err != nil {
			// error from closing listeners or context timeout
			log.Printf("Something went wrong - Server not shutting down; Error %v", err)
		}
		close(idleConnsClosed)
	}()
	// run server
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Something went wrong - Server is not running; Error %v", err)
	}
	<-idleConnsClosed
}

// simpleServer start
func SimpleServer(a *fiber.App) {
	// run it
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Something went wrong - Server is not running; Error %v", err)
	}
}
