# Hexagonal Architecture Implementation - Chat App

This repository contains the backend of a chat application implemented using Golang WebSockets, MySQL as the database, and Docker. The architecture used is the Hexagonal Architecture.

## Usage Instructions

1. Navigate to the "docker" folder using the following command:
   ```bash
   cd docker
   ```

2. From the "docker" folder, run the following command to start the Docker containers:
   ```bash
   docker-compose up -d
   ```

3. Next, to manage project dependencies:
   ```bash
   go mod tidy
   ```
4. Modify the "config.toml" file to configure the application:

    ```toml
    # Database  config
    DB_HOST = "192.168.18.135"
    DB_PORT = "3306"
    DB_USER = "root"
    DB_PASSWORD = "secret"
    DB_NAME = "chat"

    # API configuration
    ALLOWED_METHODS = ["GET", "POST", "PUT", "DELETE"]
    ALLOWED_ORIGINS = ["http://localhost:6600", "http://localhost:3000"]
    ALLOW_CREDENTIALS = true

    # Server configuration
    PORT = ":3001"
    ```
5. Finally, execute the following command to run the application:
   ```bash
    go run main.go
    ```

