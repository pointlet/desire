// pkg/config/db.go

package config

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "local"
	password = "password"
	dbname   = "farrago"
)

type DBConfig struct {
	ConnectionString string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		ConnectionString: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname),
	}
}
