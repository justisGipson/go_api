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
	// start goroutine here for concurrent execution threads
	// go routines run in the same address space
	// access to shared memory must be synchronized
	go func() {
		// buffered channels
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // catch signals from OS
		// wait for signint to send
		// data flows out, direction of arrow
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
	// wait for idleConnsClosed to send
	<-idleConnsClosed
}

// SimpleServer start up
// no goroutines or channels
func SimpleServer(a *fiber.App) {
	// run it
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Something went wrong - Server is not running; Error %v", err)
	}
}
