### Development Environment Setup

1. Clone the repository

   ```sh
   git clone https://github.com/Real-Estate-Security/backend-real-estate-ledger.git
   cd backend_real_estate
   ```

   or if you have set up SSH keys:

   ```sh
   git clone git@github.com:Real-Estate-Security/backend-real-estate-ledger.git

   ```

2. Add the school repository as a remote

   ```sh
   git remote add school https://github.com/SP25-CSCE482-capstone/backend_main_real_estate_security.git
   ```

   or if you have set up SSH keys:

   ```sh
   git remote add school git@github.com:SP25-CSCE482-capstone/backend_main_real_estate_security.git
   ```

3. Install the dependencies

   1. Install [go version 1.23](https://go.dev/doc/install)
   2. Install [Docker v20+ or latest](https://docs.docker.com/get-started/get-docker/)
   3. Install [Docker Compose](https://docs.docker.com/compose/install/)
   4. Install [go-migrate v4.18+](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) for database migrations
   5. Install [sqlc v1.28+](https://docs.sqlc.dev/en/latest/overview/install.html) for generating Go code from SQL
   6. Install [TablePlus Latest Version](https://tableplus.com/download) or [DBeaver Latest Verions](https://dbeaver.io/download/) for DB management; I prefer TablePlus
   7. Install [Postman Latest Verions](https://www.postman.com/downloads/) or [Insomnia Latest Verions](https://insomnia.rest/download) for API testing; I prefer Insomnia
   8. Install [VSCode](https://code.visualstudio.com/) or any other code editor for development

4. Docker setup and running development server

   - Ensure Docker is installed and running on your machine.
     - You can check if Docker is running by running `docker --version` in your terminal.
     - If Docker is not running, start the Docker Desktop application on your machine.
   - Run the following command to start the Docker containers:

     ```sh
        docker-compose up
     ```

5. Access the application
   - The application should now be running at `localhost:8080`.

### Additional Information: Make commands

- **Running Tests**: To run the test suite, use:

  - For all tests:

  ```sh
  make test
  ```

  - For Integration tests:

  ```sh
    make itest
  ```

- **Linting and Formatting**: To auto format the code use:

  ```sh
  make format
  ```

- **Environment Variables**: Ensure you have a `.env` file in the root of the project with the necessary environment variables. You can use the `.env.example` file as a template.

- **Database Migrations**: To run database migrations, use:

  ```sh
  make migrate
  ```

### MakeFile Commands

- You can run the makefile commands either in the docker container or just on your local machine. It might be easier to do it on your local machine. The Makefile contains the following commands:

- Run commands inside the docker container

```bash
docker exec -it dev-real-estate-backend sh
```

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container and Application Container **(DO NOT RUN IF YOU ARE INSIDE THE CONTAINER)**

```bash
make docker-run
```

Shutdown DB Container and Application Container **(DO NOT RUN IF YOU ARE INSIDE THE CONTAINER)**

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```

## Development Workflow

### Project Structure

The project is structured as follows:

```
.air.toml
.env.example
.github/
    PULL_REQUEST_TEMPLATE/
        bug_fix_template.md
        feature_template.md
    workflows/
        ...
.gitignore
.idea/
    backend_real_estate.iml
    modules.xml
    vcs.xml
    workspace.xml
app.env
cmd/
    api/
DEVELOPMENT.md
docker-compose.yml
Dockerfile
docs/
    api/
go.mod
go.sum
internal/
    database/
    server/
    token/
main
Makefile
README-Example-Final-SP25.md
README.md
Simple Real Estate ERD.png
sqlc.yml
start.sh
tmp/
util/
    config.go
    password_test.go
    password.go
    random.go
wait-for.sh
```

- **`.air.toml`**: Configuration file for [Air](https://github.com/air-verse/air)
- **`.env.example`**: Example environment variables file
- **`.github/`**: GitHub configuration files
- **`.gitignore`**: Git ignore file
- **`internal/`**: Internal packages
  - **`database/`**: Database package to handle database connections, migrations, and crud operation queries.
  - **`server/`**: Server package to handle the server setup, routes, and handlers.
  - **`token/`**: Token package to handle JWT token generation and verification.
- **`util/`**: Utility package to handle configuration, password hashing, and random string generation.
- **`cmd/`**: Command line interface for the application
  - **`api/`**: API server
- **`docs/`**: API and DB documentation
- **`go.mod`**: Go module file which lists dependencies

### Adding routes and handlers

- Read up on SQLC and how to use it to generate Go code from SQL queries.
- Add your SQL queries to a file under `internal/database/sql` and run `sqlc generate` to generate the Go code.
  - This should define the CRUD operations for your model.
- Add a handler function in a file under `internal/server` to handle the route.
  - This should call the appropriate database function.
- Add a route with the corresponding handler to the `internal/server/routes.go` file.
  - Keep the routes organized by model.
  - Use the `GET`, `POST`, `PUT`, and `DELETE` functions to define routes.
  - Use the `auth` middleware to protect routes that require authentication.

### Branching and Committing

- **Branch Naming Convention**: Use the following naming convention for your branch:

  - `feature/branch-name` for new features
  - `bugfix/branch-name` for bug fixes
  - `chore/branch-name` for maintenance tasks

- **Commit Messages**: Use the following format for your commit messages:

  ```sh
  git commit -m "type: subject"
  ```

  - The `type` can be `feat`, `fix`, `chore`, `docs`, `style`, `refactor`, or `test`.
  - The `subject` should be a short, imperative tense description of the change.

- **Pull Requests**: Use the provided templates in `.github/PULL_REQUEST_TEMPLATE` when opening a pull request. Include a clear description of the changes and the issue they address.
- **Code Reviews**: All pull requests must be reviewed by at least one other team member before merging.
