package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ashtishad/go-project-template/internal/common"
	"github.com/ashtishad/go-project-template/internal/infra/postgres"
)

func main() {
	// 1. Initialize structured logger
	handlerOpts := common.GetSlogConf()
	logger := slog.New(slog.NewTextHandler(os.Stdout, handlerOpts))
	slog.SetDefault(logger)

	// 2. Check environment variables, if not exists sets default.
	sanityCheck(logger)

	// 3. Get postgres database client
	dbClient := postgres.GetDBClient(logger)

	defer dbClient.Close()

	logger.Info("Hello from go project template")
}

// sanityCheck checks essential env variables required ot run the app, sets defaults if not exists
func sanityCheck(l *slog.Logger) {
	defaultEnvVars := map[string]string{
		"DB_USER":   "ash",
		"DB_PASSWD": "strong_password",
		"DB_HOST":   "127.0.0.1",
		"DB_PORT":   "5432",
		"DB_NAME":   "dbname",
	}

	for key, defaultValue := range defaultEnvVars {
		if os.Getenv(key) == "" {
			if err := os.Setenv(key, defaultValue); err != nil {
				l.Error(fmt.Sprintf(
					"failed to set environment variable %s to default value %s. Exiting application.",
					key,
					defaultValue,
				))
				os.Exit(1)
			}

			l.Warn(fmt.Sprintf("environment variable %s not defined. Setting to default: %s", key, defaultValue))
		}
	}
}
