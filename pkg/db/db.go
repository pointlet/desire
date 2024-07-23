package db

import (
	"context"
	"fmt"

	"github.com/feldtsen/farrago/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB(config *config.DBConfig) (*pgxpool.Pool, error) {
	// Connect to the database
	dbpool, err := pgxpool.New(context.Background(), config.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	fmt.Println("Connected to the database")
	return dbpool, nil
}
