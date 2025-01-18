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
echo "==> Building project..."
go build -o bin/recipe-manager ./cmd/recipe-manager

# 4. Test
echo "==> Running tests..."
go test -v ./...

echo "==> CI script completed successfully!"
