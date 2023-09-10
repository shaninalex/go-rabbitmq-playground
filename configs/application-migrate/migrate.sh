#!/bin/bash

# Get the PostgreSQL database credentials from environment variables
DB_HOST="$DB_HOST"
DB_PORT="$DB_PORT"
DB_NAME="$DB_NAME"
DB_USER="$DB_USER"
PGPASSWORD="$PGPASSWORD"

# Path to the SQL migration file
MIGRATION_FILE="./schema.sql"

# Construct the psql command
PSQL_COMMAND="psql -h $DB_HOST -p $DB_PORT -d $DB_NAME -U $DB_USER -W"

# Execute the migration
$PSQL_COMMAND -f $MIGRATION_FILE