package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"go-example/stores"
)

func run(ctx context.Context,
	config *Config,
	stdout io.Writer,
	stderr io.Writer,
) error {
	// allow
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// validate config
	if err := config.Validate(); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	// intitialize deps
	logger := slog.New(slog.NewJSONHandler(stdout, &slog.HandlerOptions{
		Level: config.LogLevel,
	}))

	loggingMiddleware := newLoggingMiddleware(logger)
	sessionMiddleWare := newSessionMiddleware(logger)

	pokemonStore := stores.NewInMemoryPokemonStore()

	// create server
	srv := NewServer(
		logger,
		config,
		loggingMiddleware,
		sessionMiddleWare,
		pokemonStore,
	)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: srv,
	}

	// start server
	go func() {
		logger.Info(fmt.Sprintf("listening on %s", httpServer.Addr))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(fmt.Sprintf("error listening and serving: %s", err))
			cancel()
		}
	}()

	// handle server shutdown
	var wg sync.WaitGroup
	wg.Go(func() {
		<-ctx.Done()

		logger.Info("shutting down...")
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.Error(fmt.Sprintf("error shutting down http server: %s", err))
		}
	})
	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()

	conf, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %s\n", err)
		os.Exit(1)
	}

	if err := run(ctx, conf, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
