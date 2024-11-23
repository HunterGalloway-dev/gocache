# GoCache Proof of Concept

## Overview

GoCache is a proof of concept project demonstrating a key value store implemented in Go. This project connects to an API and a database, preloading data for efficient retrieval and management. The primary goal is to showcase the integration of a key value store with a Go application, providing secure storage and access to sensitive information.

## Features

- **Key Value Store**: Securely store and manage keys using a key value store.
- **API Integration**: Connect to an API for data retrieval and manipulation.
- **Database Integration**: Preload and manage data in a database.
- **Efficient Data Retrieval**: Optimize data access and retrieval for performance.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- MongoDB or PostgreSQL (depending on your preference)
- Redis (for caching, optional)

### Installation

1. Clone the repository:
  ```sh
  git clone https://github.com/yourusername/gocache.git
  cd gocache
  ```

2. Install dependencies:
  ```sh
  go mod tidy
  ```

3. Set up your environment variables:
  ```sh
  cp .env.example .env
  ```

4. Update the `.env` file with your database and API credentials.

### Running the Application

1. Start the database and cache services (if using Docker):
  ```sh
  docker-compose up -d
  ```

2. Run the application:
  ```sh
  go run main.go
  ```

### Testing

Run the unit tests to ensure everything is working correctly:
```sh
go test ./...
```

All of these commands are simplified in the Makefile for ease of access

## Usage

### API Endpoints

- **GET /health**: Check the health of the server and database.
- **GET /data**: Retrieve preloaded data from the database.
- **POST /data**: Add new data to the database.

### Key Value Store Operations

- **Create Key**: Securely create and store a new key in the key value store.
- **Retrieve Key**: Access a stored key from the key value store.
- **Delete Key**: Remove a key from the key value store.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Acknowledgements
- Hunter Galloway (Engineer)
- MongoDB Go Driver
- Docker
- Go Documentation
- PostgreSQL Documentation
- Redis Documentation
- Docker Documentation