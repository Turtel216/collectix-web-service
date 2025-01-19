# Collections API

> ⚠️ **Disclaimer**: This project is currently a **Work in Progress (WIP)** and is not fully finished. Some features may be incomplete or unstable. Contributions and feedback are welcome as I continue to improve the application!

A web service to store and manage your collections, built with **Golang**, **Chi**, **Redis**, and containerized using **Docker** and **Docker-compose**. This API allows users to add, retrieve, update, and delete items in their geeky collections, such as comic books, action figures, trading cards, or any other collectible items.

## Features

- **CRUD operations** for managing collections and items.
- **In-memory data storage** using Redis for fast performance.
- Lightweight and efficient **REST API** built with Go and Chi.
- Containerized with Docker and easily deployable using Docker Compose.
- Clean, modular, and extensible codebase.

---

## Prerequisites

Before running the project, ensure you have the following installed:

- [Golang](https://golang.org/doc/install) (v1.20+)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/geeky-collections-api.git
cd geeky-collections-api
```

### 2. Build and Run with Docker Compose

1. Ensure Docker is running.
2. Build and run the containers:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`.

### 3. API Endpoints

| Method | Endpoint                  | Description                      |
|--------|---------------------------|----------------------------------|
| GET    | `/collections`            | List all collections             |
| POST   | `/collections`            | Create a new collection          |
| GET    | `/collections/{id}`       | Get details of a specific collection |
| PUT    | `/collections/{id}`       | Update a specific collection     |
| DELETE | `/collections/{id}`       | Delete a specific collection     |
| POST   | `/collections/{id}/items` | Add an item to a collection      |
| GET    | `/collections/{id}/items` | List items in a collection       |

### 4. Environment Variables

The application supports the following environment variables, configured in `.env`:

| Variable         | Default      | Description                           |
|------------------|--------------|---------------------------------------|
| `REDIS_HOST`     | `redis`      | Redis host name                       |
| `REDIS_PORT`     | `6379`       | Redis port                            |
| `APP_PORT`       | `8080`       | Port where the API is exposed         |

### 5. Running Tests

To run the tests, use the following command:

```bash
go test ./...
```

---

## Project Structure

```
.
├── cmd
│   ├── main.go                 # Entry point for the application
│   └── web
│       ├── app.go              # Application initialization
│       ├── collections.go      # Handlers for collection-related endpoints
│       └── routes.go           # API route definitions
├── internal
│   ├── models
│   │   └── collections.go      # Data models for collections and items
│   └── repository
│       └── collection
│           └── redis.go        # Redis-based storage implementation
├── .dockerignore               # Docker ignore file
├── .gitignore                  # Git ignore file
├── docker-compose.yml          # Docker Compose configuration
├── Dockerfile                  # Docker configuration for the API
├── go.mod                      # Go module file
├── go.sum                      # Go dependency file
├── LICENSE                     # License information
└── README.md                   # Project documentation
```

---

## Development

### Running Locally Without Docker

1. Start a Redis instance on your machine or use a cloud-hosted Redis.
2. Set the `REDIS_HOST` and `REDIS_PORT` environment variables in a `.env` file or export them in your terminal.
3. Run the application:

```bash
go run main.go
```

The API will start at `http://localhost:8080`.

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your feature/fix.
3. Commit your changes and push your branch.
4. Open a Pull Request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- Built with [Golang](https://golang.org/) and [Chi](https://github.com/go-chi/chi).
- In-memory storage powered by [Redis](https://redis.io/).
- Containerized with [Docker](https://www.docker.com/).

---

## Contact

For questions, feedback, or support, please open an issue in the repository or reach out to the maintainer at [papakonstantinou.dm@gmail.com].
