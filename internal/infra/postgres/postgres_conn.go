package postgres

import (
	"database/sql"
	"log/slog"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// GetDsnURL constructs a PostgreSQL Data Source Name *url.URL using environment variables.
// It sets the connection parameters such as user, password, host, port, database name, timezone, and SSL mode.
// The resulting DSN URL is in the format:
// "postgres://user:password@host:port/dbname?sslmode=disable&timezone=UTC"
// postgres://postgres:postgres@127.0.0.1:5432/datalake?sslmode=disable&timezone=UTC
func GetDsnURL() *url.URL {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWD")),
		Host:   net.JoinHostPort(os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Path:   os.Getenv("DB_NAME"),
	}

	q := dsn.Query()
	q.Set("timezone", "UTC")
	q.Set("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	return &dsn
}

// GetDBClient creates a new database connection and returns *sql.DB instance.
func GetDBClient(l *slog.Logger) *sql.DB {
	connConfig, err := pgx.ParseConfig(GetDsnURL().String())

	if err != nil {
		l.Error("error parsing pgx conn config", "err", err)
		os.Exit(1)
	}

	db := stdlib.OpenDB(*connConfig)

	if err = db.Ping(); err != nil {
		l.Error("error pinging the database", "err", err.Error())
		os.Exit(1)
	}

	l.Info("successfully connected to database", "dsn", connConfig.ConnString())

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
