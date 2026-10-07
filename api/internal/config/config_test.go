package config_test

import (
	"context"
	"testing"
	"time"

	"github.com/davidporos92/margin-cms/api/internal/config"
	"github.com/davidporos92/margin-cms/api/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestNew_Defaults(t *testing.T) {
	expectedCfg := &config.Config{
		Server: &server.Config{
			Addr:              ":8080",
			ShutdownTimeout:   30 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
		},
	}

	assert.Equal(t, expectedCfg, config.New(context.TODO()))
}
