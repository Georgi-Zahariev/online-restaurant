#!/bin/bash

export PORT=${PORT:-8080}

echo "Starting server on port $PORT..."
cd backend
go run main.go
