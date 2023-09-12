#!/bin/bash

# Get the PostgreSQL database credentials from environment variables
DB_HOST="$POSTGRES_HOST"
DB_PORT=5432
DB_NAME="$POSTGRES_DB"
DB_USER="$POSTGRES_USER"
PGPASSWORD="$POSTGRES_PASSWORD"

# Path to the SQL migration file
MIGRATION_FILE="./schema.sql"

# Construct the psql command
PSQL_COMMAND="psql -h $DB_HOST -p $DB_PORT -d $DB_NAME -U $DB_USER -W"

# Execute the migration
$PSQL_COMMAND -f $MIGRATION_FILE