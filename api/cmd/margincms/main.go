package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/davidporos92/margin-cms/api/internal/api"
	"github.com/davidporos92/margin-cms/api/internal/config"
	"github.com/davidporos92/margin-cms/api/internal/server"
	"github.com/go-chi/cors"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.New(ctx)

	serverInterface := api.NewStrictHandlerWithOptions(server.New(), server.NewServerMiddlewares(), server.NewServerOptions())
	handler := api.HandlerWithOptions(serverInterface, api.ChiServerOptions{
		BaseURL: "/api/v1",
	})
	handler = cors.Handler(cors.Options{
		AllowedOrigins:   cfg.Server.AllowedOrigins,
		AllowedHeaders:   cfg.Server.AllowedHeaders,
		AllowedMethods:   cfg.Server.AllowedMethods,
		AllowCredentials: cfg.Server.AllowCredentials,
	})(handler)
	srv := &http.Server{
		Addr:              cfg.Server.Addr,
		Handler:           handler,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		ReadTimeout:       cfg.Server.ReadTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}
