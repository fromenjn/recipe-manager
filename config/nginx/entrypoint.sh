#!/bin/sh

# Start the Go app in the background (listening on port 8080)
# Make sure your Go program listens on 0.0.0.0:8080 internally
cd /app && ./recipe-manager &

# Run Nginx in the foreground (so the container doesn't exit)
exec nginx -g "daemon off;"