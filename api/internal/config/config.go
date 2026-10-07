package config

import (
	"context"
	"log"

	"github.com/davidporos92/margin-cms/api/internal/server"
	"github.com/sethvargo/go-envconfig/v2"
)

type Config struct {
	Server *server.Config
}

func New(ctx context.Context) *Config {
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	return &c
}
