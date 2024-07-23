package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type UserAccount struct {
	ID           int
	Username     string
	PasswordHash string
}

func GetUserAccountEntry(db *pgx.Conn, username string) (*UserAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, username, password_hash 
		FROM user_accounts 
		WHERE username = $1
	`

	row := db.QueryRow(ctx, query, username)

	userAccount := &UserAccount{}

	err := row.Scan(&userAccount.ID, &userAccount.Username, &userAccount.PasswordHash)
	if err != nil {
		return nil, err
	}

	return userAccount, nil
}
