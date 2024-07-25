# Go APP
This is a simple Go application

## Prerequisites

- Docker (https://www.docker.com/get-started)

## Installation

Clone the repository

## Running the Application

1. Run Docker Compose:

    ```bash
    docker-compose up
    ```

2. copy ./schema/migrations/migrate.sql into your database to create new table

3. Run the application:

    ```bash
    go run . rest
    ```

## Testing the Application
import ./additionaldata/postman collection.json to your postman and try to send data