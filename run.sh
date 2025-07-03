#!/bin/bash

# Set environment variables with defaults if not already set
export PORT=${PORT:-8080}
export ENV=${ENV:-development}
export LOG_FORMAT=${LOG_FORMAT:-text}
export LOG_SEVERITY=${LOG_SEVERITY:-debug}

# DB envs for local runs (override in your shell or CI):
export DB_HOST="${DB_HOST:-localhost}"
export DB_PORT="${DB_PORT:-5432}"
export DB_NAME="${DB_NAME:-restaurant}"
export DB_USER="${DB_USER:-postgres}"
export DB_PASSWORD="${DB_PASSWORD:-postgres}"
export DB_SSLMODE="${DB_SSLMODE:-disable}"

DB_CONTAINER_NAME="restaurant_db"

# Check if the DB container is running, start it if not
if ! docker ps --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$"; then
  echo "Database container '${DB_CONTAINER_NAME}' is not running. Starting it..."
  if ! docker ps -a --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$"; then
    docker run -d \
      --name ${DB_CONTAINER_NAME} \
      -e POSTGRES_DB=${DB_NAME} \
      -e POSTGRES_USER=${DB_USER} \
      -e POSTGRES_PASSWORD=${DB_PASSWORD} \
      -p ${DB_PORT}:5432 \
      postgres:17
  else
    docker start ${DB_CONTAINER_NAME}
  fi
else
  echo "Database container '${DB_CONTAINER_NAME}' is already running."
fi

# Wait for DB to be ready
echo "Waiting for database to be ready..."
for i in {1..10}; do
  if docker exec ${DB_CONTAINER_NAME} pg_isready -U "${DB_USER}" -d "${DB_NAME}" > /dev/null 2>&1; then
    echo "Database is ready!"
    break
  fi
  sleep 1
done

echo "Starting backend..."
echo "  PORT:         $PORT"
echo "  ENV:          $ENV"
echo "  LOG_FORMAT:   $LOG_FORMAT"
echo "  LOG_SEVERITY: $LOG_SEVERITY"
echo " DB: $DB_USER@$DB_HOST:$DB_PORT/$DB_NAME (sslmode=$DB_SSLMODE)"

cd backend
go run .