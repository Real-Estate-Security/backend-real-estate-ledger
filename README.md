# Real Estate Transaction Ledger Backend

## Project Title and Description

The **Real Estate Transaction Ledger Backend** is a Go Gin server that implements REST API routes for a real estate transaction ledger. The backend communicates with a Hyperledger Fabric network to commit transaction data, enabling a secure and transparent system for managing real estate transactions.

## Archetecture and Design

- [Database Documentation](https://dbdocs.io/malikrawashdeh/real_estate_ledger_db)
- [API Documentation] : Coming Soon
- System Architecture:

## Table of Contents

- [Real Estate Transaction Ledger Backend](#real-estate-transaction-ledger-backend)
  - [Project Title and Description](#project-title-and-description)
  - [Archetecture and Design](#archetecture-and-design)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
    - [Environment](#environment)
    - [Program](#program)
    - [Tools](#tools)
  - [External Dependencies](#external-dependencies)
  - [Environmental Variables/Files](#environmental-variablesfiles)
  - [Installation and Setup](#installation-and-setup)
    - [Clone the Repository](#clone-the-repository)
    - [Install Dependencies](#install-dependencies)
    - [Set Up Docker and Run the Development Server](#set-up-docker-and-run-the-development-server)
  - [Development Workflow](#development-workflow)
    - [Running Tests](#running-tests)
    - [Linting and Formatting](#linting-and-formatting)
    - [Database Migrations](#database-migrations)
    - [Makefile Commands](#makefile-commands)
  - [Usage](#usage)
  - [Features](#features)
  - [Documentation](#documentation)
  - [Credits and Acknowledgments](#credits-and-acknowledgments)
  - [License](#license)
  - [Third-Party Libraries](#third-party-libraries)
  - [Contact Information](#contact-information)

## Requirements

This code has been run and tested using the following internal and external components:

### Environment

- Docker Engine v20+
- Docker Compose v2.13+
- Go v1.23

### Program

- PostgreSQL v17 or later
- sqlc v1.28+
- go-migrate v4.18+
- TablePlus or DBeaver for database management
- Postman or Insomnia for API testing

### Tools

- GitHub - main branch with [repository link](https://github.com/Real-Estate-Security/backend-real-estate-ledger)
- VSCode or any other code editor

## External Dependencies

- [Docker](https://www.docker.com/products/docker-desktop)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go](https://go.dev/doc/install)

## Environmental Variables/Files

Ensure you have a `.env` file in the root of the project with the necessary environment variables. You can use the `.env.example` file as a template.

## Installation and Setup

### Clone the Repository

```sh
git clone https://github.com/Real-Estate-Security/backend-real-estate-ledger.git
cd backend_real_estate
```

or if you have set up SSH keys:

```sh
git clone git@github.com:Real-Estate-Security/backend-real-estate-ledger.git
```

### Install Dependencies

1. Install Go v1.23
2. Install Docker v20+
3. Install Docker Compose
4. Install go-migrate v4.18+
5. Install sqlc v1.28+
6. Install a database management tool (TablePlus/DBeaver)
7. Install an API testing tool (Postman/Insomnia)

### Set Up Docker and Run the Development Server

- Ensure Docker is installed and running:
  ```sh
  docker --version
  ```
- Start the Docker containers:
  ```sh
  docker-compose up
  ```
- The application should now be running at `localhost:8080`

## Development Workflow

### Running Tests

- Run all tests:
  ```sh
  make test
  ```
- Run integration tests:
  ```sh
  make itest
  ```

### Linting and Formatting

- Auto format code:
  ```sh
  make format
  ```

### Database Migrations

- Run database migrations:
  ```sh
  make migrate
  ```

### Makefile Commands

- Build and run the application:
  ```sh
  make all
  make build
  make run
  ```
- Create and manage Docker containers:
  ```sh
  make docker-run  # Start DB and application container
  make docker-down # Stop DB and application container
  ```
- Live reload the application:
  ```sh
  make watch
  ```
- Clean up binary from the last build:
  ```sh
  make clean
  ```

## Usage

The backend server exposes REST API routes that the frontend can call to interact with the real estate transaction ledger. The API communicates with a Hyperledger Fabric network to commit transaction data securely.

## Features

- Secure transaction ledger for real estate transactions
- REST API implementation using Go Gin
- Integration with Hyperledger Fabric for transaction data
- Dockerized development environment
- Automated database migrations
- CI/CD-ready with Makefile commands

## Documentation

For additional documentation and API references, visit the `/docs/api/` directory in the repository.

## Credits and Acknowledgments

- [Real Estate Security Team](https://github.com/Real-Estate-Security) for development and maintenance

## License

For now this project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Third-Party Libraries

- Go Gin (MIT License)
- sqlc (MIT License)
- go-migrate (MIT License)

## Contact Information

For any questions, contact:

- [GitHub Issues](https://github.com/Real-Estate-Security/backend-real-estate-ledger/issues)
- Contributors:
  - @malikrawashdeh : Malik Rawashdeh
  - @cj24787 : Caroline Jia
  - @mahikaperi : Mahika Peri
  - @saraha245 : Sarah Ahmed
  - @briancoco : Brian Nguyen
