#!/bin/bash

# Set environment variables with defaults if not already set
export PORT=${PORT:-8080}
export ENV=${ENV:-development}
export LOG_FORMAT=${LOG_FORMAT:-text}     # options: text or json
export LOG_SEVERITY=${LOG_SEVERITY:-debug} # options: debug, info, warn, error

echo "Starting backend..."
echo "  PORT:         $PORT"
echo "  ENV:          $ENV"
echo "  LOG_FORMAT:   $LOG_FORMAT"
echo "  LOG_SEVERITY: $LOG_SEVERITY"

cd backend
go run .

