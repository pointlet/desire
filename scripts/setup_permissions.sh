#!/bin/bash

# Define variables
SUPERUSER="farrago"   # replace with actual superuser
PASSWORD="password"    # replace with actual password
DATABASE="farrago"
SQL_FILE="grant_permissions.sql"

# Export password to avoid password prompt
export PGPASSWORD=$PASSWORD

# Execute the SQL script as the superuser
psql -U $SUPERUSER -d $DATABASE -f $SQL_FILE

# Unset the password variable for security
unset PGPASSWORD

# Feedback to the user
echo "Permissions granted to the 'local' user."

