package main

import (
	"fmt"
	"log/slog"

	"github.com/mist-gopher/feature-tag/configs"
	"github.com/mist-gopher/feature-tag/pkg/logger"
)

func main() {
	config := configs.Load()
	logger.Configure(config.LogLevel)

	greeting := "Hello"
	target := "World"
	message := fmt.Sprintf("%s %s!", greeting, target)

	slog.Info(message)
}
