package internal

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutdownDelay = 5 * time.Second

func gracefulShutdown(server *fiber.App, serverCloseChan chan string) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)

	// Makes the goroutine wait for a message to be sent into the channel
	<-sigint

	log.Println("Stopping Server")
	time.Sleep(shutdownDelay)

	_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := server.Shutdown(); err != nil {
		log.Fatalf("Server shutdown error: %v\n", err)
	}

	serverCloseChan <- "Server shutdown complete"
	close(serverCloseChan)
}
