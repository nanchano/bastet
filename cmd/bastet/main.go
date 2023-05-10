package main

import (
	"os"

	"github.com/nanchano/bastet/internal/config"
	"github.com/nanchano/bastet/internal/core"
	"github.com/nanchano/bastet/internal/repository"
	"github.com/nanchano/bastet/internal/server"
	"golang.org/x/exp/slog"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout)
	logger := slog.New(handler)

	logger.Info("Reading config")
	config, err := config.Load(".env")
	if err != nil {
		logger.Error("Failed parsing config: %v", err)
		panic(err)
	}

	repo, err := repository.New(config.Database.URL())
	if err != nil {
		logger.Error(" Failed connecting to the database: $v", err)
		panic(err)
	}

	service := core.NewService(logger, repo)
	logger.Info("Starting servers")

	server := server.New(logger, service)
	server.Start()
}
