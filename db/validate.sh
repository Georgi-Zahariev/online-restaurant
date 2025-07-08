#!/usr/bin/env bash
set -e

# Load environment variables from .env
set -a
[ -f "$PWD/../.env" ] && source "$PWD/../.env"
set +a

ROOT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

POSTGRES_CONTAINER="restaurant-db-validate"
POSTGRES_VERSION="17-alpine"

# Use a non-standard port for validation to avoid conflicts, or fallback to .env port
VALIDATE_PORT=55432
POSTGRES_PORT="${VALIDATE_PORT}"

CONNECTION_STRING="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}"

function clean() {
    echo "Removing DB container..."
    docker rm --force ${POSTGRES_CONTAINER} >/dev/null 2>&1 || true
}
trap clean EXIT

echo "Starting fresh DB container for validation..."
docker run -d --name ${POSTGRES_CONTAINER} \
    -e POSTGRES_DB=${POSTGRES_DB} \
    -e POSTGRES_USER=${POSTGRES_USER} \
    -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
    -p ${POSTGRES_PORT}:5432 \
    postgres:${POSTGRES_VERSION}

echo "Waiting for DB to become ready..."
for i in {1..10}; do
    docker exec ${POSTGRES_CONTAINER} pg_isready -U "${POSTGRES_USER}" -h "localhost" -p "5432" -d "${POSTGRES_DB}" && break
    sleep 1
done

echo "Running migrations UP..."
migrate -path db/migrations -database "${CONNECTION_STRING}" up

if [ -d "${ROOT_PATH}/db/init" ]; then
    echo "Loading initial SQL data..."
    for f in ${ROOT_PATH}/db/init/*.sql; do
        [ -e "$f" ] || continue
        cat "$f" | docker exec -i ${POSTGRES_CONTAINER} psql -U "${POSTGRES_USER}" -d "${POSTGRES_DB}"
    done
fi

echo "Migration version after UP: $(migrate -path db/migrations -database "${CONNECTION_STRING}" version 2>&1)"

echo "Running migrations DOWN to 0..."
migrate -path db/migrations -database "${CONNECTION_STRING}" down 0

echo "Migration version after DOWN: $(migrate -path db/migrations -database "${CONNECTION_STRING}" version 2>&1)"

echo "Validation complete. Cleaning up."