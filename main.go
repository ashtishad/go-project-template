package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/ashtishad/go-project-template/internal/common"
	"github.com/ashtishad/go-project-template/internal/infra/postgres"
	"github.com/ashtishad/go-project-template/internal/infra/transport"
)

func main() {
	// initialize structured logger.
	handlerOpts := common.GetSlogConf()
	logger := slog.New(slog.NewTextHandler(os.Stdout, handlerOpts))
	slog.SetDefault(logger)

	// check environment variables, if not exists sets default.
	sanityCheck(logger)

	// get postgres database client.
	dbClient := postgres.GetDBClient(logger)

	defer dbClient.Close()

	// setting up the routers.
	router := http.NewServeMux()
	router.HandleFunc("GET /", transport.NameHandler)
	router.HandleFunc("GET /{name}", transport.NameHandler)

	// creating a custom Server.
	s := &http.Server{
		Addr:           net.JoinHostPort("", os.Getenv("API_PORT")),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// start the Server.
	logger.Info("Server starting...", slog.String("address", s.Addr))

	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("error starting server", "err", err)
		return
	}
}

// sanityCheck checks essential env variables required ot run the app, sets defaults if not exists
func sanityCheck(l *slog.Logger) {
	defaultEnvVars := map[string]string{
		"API_PORT":  "8000",
		"DB_USER":   "ash",
		"DB_PASSWD": "strong_password",
		"DB_HOST":   "127.0.0.1",
		"DB_PORT":   "5432",
		"DB_NAME":   "datalake",
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
