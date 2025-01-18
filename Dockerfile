
FROM node:22 AS frontend-builder
WORKDIR /app
COPY . .
RUN cd frontend/boca-recettes && npm install && npm run build


FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add --no-cache git bash golangci-lint
RUN go build -o bin/recipe-manager ./cmd/recipe-manager


FROM nginx:1.27.3-alpine
RUN apk --no-cache add ca-certificates tzdata

COPY config /app/config
COPY data /app/data

COPY config/nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/bin/recipe-manager /app/recipe-manager
COPY --from=frontend-builder /app/frontend/boca-recettes/dist /app/dist

# Copy our entrypoint script
COPY config/nginx/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080

# Set the entrypoint to run both the Go app and Nginx
ENTRYPOINT ["/entrypoint.sh"]