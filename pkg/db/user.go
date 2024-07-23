package db

import (
	"context"
	"time"

	"github.com/feldtsen/farrago/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUserAccountEntry(username string) (*UserAccount, error)
	InsertUserAccountEntry(username, passwordHash string) (models.DataManipulationResult, error)
	DeleteUserAccountEntry(username string) (models.DataManipulationResult, error)
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

func (repo *PgxUserRepository) InsertUserAccountEntry(username, password string) (models.DataManipulationResult, error) {
	passwordHash, err := GenerateHashPassword(password)
	if err != nil {
		return models.DataManipulationResult{RowsAffected: 0}, err
	}

	query := `
		INSERT INTO user_accounts (username, password_hash)
		VALUES ($1, $2)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := repo.DB.Exec(ctx, query, username, passwordHash)
	if err != nil {
		return models.DataManipulationResult{RowsAffected: 0}, err
	}

	return models.DataManipulationResult{RowsAffected: result.RowsAffected()}, nil
}

func (repo *PgxUserRepository) DeleteUserAccountEntry(username string) (models.DataManipulationResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		DELETE FROM user_accounts
		WHERE username = $1 
	`

	result, err := repo.DB.Exec(ctx, query, username)
	if err != nil {
		return models.DataManipulationResult{RowsAffected: 0}, err
	}

	return models.DataManipulationResult{RowsAffected: result.RowsAffected()}, nil
}
