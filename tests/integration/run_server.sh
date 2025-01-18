#!/usr/bin/env bash
set -euo pipefail

cd ../../ && \
go build -o bin/recipe-manager ./cmd/recipe-manager && \
./bin/recipe-manager