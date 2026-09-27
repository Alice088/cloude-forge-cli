package main

import (
	"Alice088/cloude-forge-cli/cmd"
	"Alice088/cloude-forge-cli/internal/config"
	"Alice088/cloude-forge-cli/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const version = "0.1.0"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("failed to .env")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	logger.Setup(cfg.Log)

	cmd.Execute()
}
