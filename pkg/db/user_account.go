package db

import "database/sql"

/*
CREATE TABLE IF NOT EXISTS user_accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(60) NOT NULL,
    -- last_login TIMESTAMP WITH TIME ZONE,
    -- failed_login_attempts INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- User Profile Table
CREATE TABLE IF NOT EXISTS user_profiles (
    id SERIAL PRIMARY KEY,
    user_account_id INTEGER UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE,
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE
);

-- Indexes for faster lookups
CREATE INDEX IF NOT EXISTS idx_user_accounts_username ON user_accounts(username);
CREATE INDEX IF NOT EXISTS idx_user_profiles_email ON user_profiles(email);

Above are my tables, help me create a select to get the hashed_password for a user with a specific username.
*/

func GetUserAccountEntry(db *sql.DB, username string) (string, error) {
	var passwordHash string

	// Directly execute the query using QueryRow
	err := db.QueryRow("SELECT password_hash FROM user_accounts WHERE username = $1", username).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // return an error indicating no user found
		}
		// Return any other error that occurred during the query execution
		return "", err
	}

	// Return the password hash
	return passwordHash, nil
}

func SetUserAccountEntry(db *sql.DB, username, passwordHash string) error {
	// Prepare the query
	query := "INSERT INTO user_accounts (username, password_hash) VALUES ($1, $2)"
	// Execute the query
	_, err := db.Exec(query, username, passwordHash)
	if err != nil {
		// Return the error if the query execution failed
		return err
	}

	// Return nil if the operation was successful
	return nil
}
