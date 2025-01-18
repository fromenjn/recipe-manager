#!/usr/bin/env bash
set -euo pipefail

# 1. Lint
echo "==> Running linters..."
golangci-lint run ./...

# 2. Generate Swagger documentation
echo "==> Generating Swagger docs..."
swag init \
  -d cmd/recipe-manager \
  -o docs -pdl 3

# 3. Build
echo "==> Building backend..."
go build -o bin/recipe-manager ./cmd/recipe-manager

# 4. Test
echo "==> Running tests..."
go test -v ./...

echo "==> Backend CI script completed successfully!"


# Frontend
echo "==> Building frontend..."
rm -rf ./dist
cd ./frontend/boca-recettes/ && npm run build && mv dist ../../
