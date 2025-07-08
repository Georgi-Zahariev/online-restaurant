#!/bin/bash

# export all varaibles in .ev file to here
set -a
[ -f .env ] && source .env
set +a

DB_CONTAINER_NAME="restaurant_db"

# keepDBs variable: if set, reuse DB containers; if not, remove and recreate
if [ -z "$keepDBs" ]; then
  echo "keepDBs not set. Removing any existing DB containers..."
  if docker ps -a --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$"; then
    docker rm --force ${DB_CONTAINER_NAME}
  fi
else
  echo "keepDBs is set. Will reuse existing DB containers if present."
fi

# Check if the DB container is running, start it if not
if ! docker ps --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$"; then
  echo "Database container '${DB_CONTAINER_NAME}' is not running. Starting it..."
  if ! docker ps -a --format '{{.Names}}' | grep -q "^${DB_CONTAINER_NAME}$"; then
    docker run -d \
      --name ${DB_CONTAINER_NAME} \
      -e POSTGRES_DB=${POSTGRES_DB} \
      -e POSTGRES_USER=${POSTGRES_USER} \
      -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
      -p ${POSTGRES_PORT}:5432 \
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
  if docker exec ${DB_CONTAINER_NAME} pg_isready -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" > /dev/null 2>&1; then
    echo "Database is ready!"
    break
  fi
  sleep 1
done

echo "Starting backend..."
echo