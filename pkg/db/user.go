package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUserAccountEntry(username string) (*UserAccount, error)
}

type UserAccount struct {
	ID           int
	Username     string
	PasswordHash string
}

type PgxUserRepository struct {
	DB *pgxpool.Pool
}

func (repo *PgxUserRepository) GetUserAccountEntry(username string) (*UserAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, username, password_hash 
		FROM user_accounts 
		WHERE username = $1
	`

	row := repo.DB.QueryRow(ctx, query, username)

	userAccount := &UserAccount{}

	err := row.Scan(&userAccount.ID, &userAccount.Username, &userAccount.PasswordHash)
	if err != nil {
		return nil, err
	}

	return userAccount, nil
}

func (repo *PgxUserRepository) InsertUserAccountEntry(username, passwordHash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO user_accounts (username, password_hash)
		VALUES ($1, $2)
	`

	_, err := repo.DB.Exec(ctx, query, username, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PgxUserRepository) DeleteUserAccountEntry(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		DELETE FROM user_accounts
		WHERE username = $1 
	`

	_, err := repo.DB.Exec(ctx, query, username)
	if err != nil {
		return err
	}

	return nil
}
