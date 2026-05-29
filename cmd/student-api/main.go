package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/satyamkanungo-dev/student-api/internal/config"
)

func main() {
	// config setup
	cfg := config.MustLoad()
	// database setup

	// router setup
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student-api"))
	})

	// server setup
	// using go routines and channel

	// empty channel
	done := make(chan os.Signal, 1)

	svr := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address:", cfg.Addr))

	// checkout syscall const (what are they)
	// push into a channel
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// gorutine
	go func() {
		if err := svr.ListenAndServe(); err != nil {
			log.Fatal("failed to start server")
		}
	}()

	// if channel is empty, it wait (it's like a gateway)
	<-done

	// server gracefully shutdown

	// checkout slog package
	slog.Info("Shutting down the server")

	// checkout context package
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}
