package postgre

import (
	"context"
	"fmt"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"os"
	"time"
)

const (
	PGUSER     = "PGUSER"
	PGPASSWORD = "PGPASSWORD"
	PGHOST     = "PGHOST"
	PGPORT     = "PGPORT"
	PGDATABASE = "PGDATABASE"
)

var dbPool *pgxpool.Pool // Variável global para o pool de conexões

func ConnectDB() (*pgxpool.Pool, error) {
	logger.Info("Init ConnectDB", zap.String("journey", "Init Database"))

	connStr, err := getConnectionString()
	if err != nil {
		logger.Error("Failed to build connection string", err)
		return nil, fmt.Errorf("failed to build connection string: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		logger.Error("Failed to create connection pool", err)
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		logger.Error("Ping failed", err)
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	logger.Info("Connected to PostgreSQL", zap.String("journey", "Init Database"))

	dbPool = pool // Salva o pool na variável global
	return pool, nil
}

func GetDB() *pgxpool.Pool {
	return dbPool
}

func getConnectionString() (string, error) {
	envVariables := []string{PGUSER, PGPASSWORD, PGHOST, PGPORT, PGDATABASE}

	for _, envVariable := range envVariables {
		if os.Getenv(envVariable) == "" {
			return "", fmt.Errorf("There is no %v variable", envVariable)
		}
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv(PGUSER),
		os.Getenv(PGPASSWORD),
		os.Getenv(PGHOST),
		os.Getenv(PGPORT),
		os.Getenv(PGDATABASE),
	)

	return connStr, nil
}
