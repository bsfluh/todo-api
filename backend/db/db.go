package db

import (
	"context"
	"fmt"
	"todo-Api/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("conection error: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}
	fmt.Println("DataBase successfully connected")
	return pool, err
}
