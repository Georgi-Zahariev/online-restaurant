#!/bin/bash

# Set environment variables with defaults if not already set
export PORT=${PORT:-8080}
export ENV=${ENV:-development}
export LOG_FORMAT=${LOG_FORMAT:-text}     # options: text or json
export LOG_SEVERITY=${LOG_SEVERITY:-debug} # options: debug, info, warn, error

# DB envs for local runs (override in your shell or CI):
export DB_HOST     = "${DB_HOST:-localhost}"
export DB_PORT     = "${DB_PORT:-5432}"
export DB_NAME     = "${DB_NAME:-restaurant}"
export DB_USER     = "${DB_USER:-postgres}"
export DB_PASSWORD = "${DB_PASSWORD:-postgres}"
export DB_SSLMODE  = "${DB_SSLMODE:-disable}"

echo "Starting backend..."
echo "  PORT:         $PORT"
echo "  ENV:          $ENV"
echo "  LOG_FORMAT:   $LOG_FORMAT"
echo "  LOG_SEVERITY: $LOG_SEVERITY"
echo " DB: $DB_USER@$DB_HOST:$DB_PORT/$DB_NAME (sslmode=$DB_SSLMODE)"

cd backend
go run .

