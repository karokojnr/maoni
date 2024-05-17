# Maoni

Maoni is a Go API template that follows the clean architecture principles. It includes all CRUD operations.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Local Development](#local-development)
- [API Endpoints](#api-endpoints)
- [Authorization](#authorization)
- [Contributing](#contributing)

## Prerequisites

Before you begin, ensure you have installed the following:

- [Go](https://go.dev/doc/install) version 1.20+
- [Docker](https://docs.docker.com/)
- [Task](https://taskfile.dev/installation/) - Helps in running multiple commands all at once.
- [golangci-lint](https://golangci-lint.run/welcome/install/) - Helps in linting the API.

## Local Development

### Running the app

Clone the repository and run the app using Task:

```bash
git clone https://github.com/karokojnr/maoni.git
cd maoni
task run
```

### Linting the app

Lint the app using the following command:

```bash
task lint
```

### Tests

Run integration tests with:

```bash
task integration-test

```

Run end-to-end tests with:

```bash
task e2e-test

```

## API Endpoints

| Method | Endpoints           | Action                   |
| ------ | ------------------- | ------------------------ |
| POST   | /api/v1/comment     | Create a new comment     |
| GET    | /api/v1/comment/:id | Retrieve a comment by id |
| PUT    | /api/v1/comment/:id | Update a comment         |
| DELETE | /api/v1/comment/:id | Delete a comment         |

## Authorization

Some requests require a JWT Token. You can generate a JWT token at [jwt.io](jwt.io) using `c2VjcmV0ignvbnNlY3JldCJ9` as the secret.

To authenticate an API request, provide your JWT token in the `Authorization` header:

```bash
Authorization [token]
```

These requests require authentication:

    - POST   /api/v1/comment
    - PUT    /api/v1/comment/:id
    - DELETE /api/v1/comment/:id

## Contributing

If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.
