package server

import "time"

type Config struct {
	Addr              string        `env:"SERVER_ADDR, default=:8080"`
	ShutdownTimeout   time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT, default=30s"`
	ReadHeaderTimeout time.Duration `env:"SERVER_READ_HEADER_TIMEOUT, default=5s"`
	ReadTimeout       time.Duration `env:"SERVER_READ_TIMEOUT, default=30s"`
	WriteTimeout      time.Duration `env:"SERVER_WRITE_TIMEOUT, default=30s"`
}
