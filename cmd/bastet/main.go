package main

import (
	"log/slog"
	"os"

	"github.com/nanchano/bastet/internal/config"
	"github.com/nanchano/bastet/internal/core"
	"github.com/nanchano/bastet/internal/repository"
	"github.com/nanchano/bastet/internal/server"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)

	logger.Info("Reading config")
	config, err := config.Load(".env")
	if err != nil {
		logger.Error("Failed parsing config", "error", err)
		panic(err)
	}

	repo, err := repository.New(config.Database.URL())
	if err != nil {
		logger.Error("Failed connecting to the database", "error", err)
		panic(err)
	}

	service := core.NewService(logger, repo)
	logger.Info("Starting servers")

	server := server.New(logger, service)
	server.Start(config.Server.Host, config.Server.Port)
}
