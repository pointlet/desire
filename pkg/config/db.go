// pkg/config/db.go

package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "local"
	password = "password"
	dbname   = "farrago"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	// Create a connection string
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the database
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	fmt.Println("Connected to the database")
	return dbpool, nil
}
