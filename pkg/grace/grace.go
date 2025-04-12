package grace

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Serve will run HTTP server with graceful shutdown capability
func Serve(port string, h http.Handler) error {
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "10000"
		}
	}

	// Safely add leading colon if missing
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      h,
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		log.Println("🛑 Received shutdown signal")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("🚀 HTTP server running on", port)
	if err := server.Serve(lis); err != http.ErrServerClosed {
		return err
	}

	<-idleConnsClosed
	log.Println("✅ HTTP server shutdown gracefully")
	return nil
}

