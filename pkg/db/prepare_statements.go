package db

import (
	"database/sql"
	"log"
)

type PreparedStatements struct {
	GetUserAccountPasswordHash *sql.Stmt
	InsertNewUser              *sql.Stmt
}

func PrepareStatements(db *sql.DB) *PreparedStatements {
	var preparedStatements PreparedStatements
	var err error

	preparedStatements.GetUserAccountPasswordHash, err = db.Prepare("SELECT password_hash FROM user_accounts WHERE username = $1")
	if err != nil {
		log.Fatalf("Error preparing GetUserAccountEntry statement: %q", err)
	}

	preparedStatements.InsertNewUser, err = db.Prepare("INSERT INTO user_accounts (username, password_hash) VALUES ($1, $2)")
	if err != nil {
		log.Fatalf("Error preparing SetUserAccountEntry statement: %q", err)
	}

	return &preparedStatements
}

func (preparedStatements *PreparedStatements) Close() {
	if preparedStatements.GetUserAccountPasswordHash != nil {
		preparedStatements.GetUserAccountPasswordHash.Close()
	}
	if preparedStatements.InsertNewUser != nil {
		preparedStatements.InsertNewUser.Close()
	}
}
