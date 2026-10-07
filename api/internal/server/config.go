package server

import "time"

type Config struct {
	Addr              string        `env:"SERVER_ADDR, default=:8080"`
	ShutdownTimeout   time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT, default=30s"`
	AllowedOrigins    []string      `env:"SERVER_ALLOWED_ORIGINS, default=http://localhost:8081"`
	AllowedMethods    []string      `env:"SERVER_ALLOWED_METHODS, default=GET,POST,PUT,PATCH,DELETE,OPTIONS"`
	AllowedHeaders    []string      `env:"SERVER_ALLOWED_HEADERS, default=Authorization,Content-Type"`
	AllowCredentials  bool          `env:"SERVER_ALLOW_CREDENTIALS, default=false"`
	ReadHeaderTimeout time.Duration `env:"SERVER_READ_HEADER_TIMEOUT, default=5s"`
	ReadTimeout       time.Duration `env:"SERVER_READ_TIMEOUT, default=30s"`
	WriteTimeout      time.Duration `env:"SERVER_WRITE_TIMEOUT, default=30s"`
}
