#!/bin/bash

# Variables
DB_USER="local"
DB_PASSWORD="password"
DB_NAME="farrago"

# Ensure PostgreSQL is running
brew services start postgresql@16

# Create user if it doesn't exist
psql postgres -c "SELECT 1 FROM pg_roles WHERE rolname='$DB_USER'" | grep -q 1 || \
psql postgres -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD' CREATEDB;"

# Create database if it doesn't exist
psql postgres -c "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'" | grep -q 1 || \
createdb -U $DB_USER $DB_NAME

echo "Setup complete. You can now connect to your database using:"
echo "psql -U $DB_USER -d $DB_NAME"
