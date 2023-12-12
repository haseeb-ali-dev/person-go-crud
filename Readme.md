# Go CRUD (Create, Read, Update, Delete) Project

A simple Go project demonstrating CRUD operations with a RESTful API for managing person data.

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Docker Installation](#docker-installation)
- [API Documentation](#api-documentation)
- [Testing](#testing)

## Introduction

This Go CRUD project showcases a basic implementation of a RESTful API for managing person data. It includes endpoints for creating, reading, updating, and deleting person records.

## Project Structure

The project is organized as follows:

```bash
go-crud/
|-- main.go
|-- main_test.go
|-- go.mod
|-- .dockerignore
|-- Dockerfile
|-- compose.yml
|-- handlers/
|   |-- health.go
|   |-- person.go
|-- data/
|   |-- person.go
|-- collection/
|   |-- thunder-collection_Person Data API -- Go.json
```

<br>

- `main.go`: The main entry point of the application.
- `main_test.go`: Test cases for the project.
- `go.mod`: Go module file for managing dependencies.
- `handlers/`: Contains HTTP handlers for different endpoints.
- `data/`: Manages the data layer including CRUD operations.
- `collection/`: Holds API documentation in Thunder Collection format.
- `.dockerignore`: Specifies files and directories to be excluded during Docker builds.
- `Dockerfile`: Defines how to build the Docker image.
- `compose.yml`: Orchestrates the application using Docker Compose.

## Installation

To install and run the project, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/haseeb-ali-dev/go-crud.git
   ```

2. Change into the project directory:

   ```bash
   cd go-crud
   ```

3. Install dependencies:

   ```go
   go mod download
   ```

## Usage

To run the project, execute the following command:

```go
go run main.go
```

The server will start on **`localhost:4444`**. You can now interact with the API using your preferred API client.

## Docker Installation

1. Make sure you have Docker installed on your machine. If not, follow the official instructions for [Docker installation](https://docs.docker.com/get-docker/).

2. To run the project using Docker, execute the following commands:

   - Build the Docker image:

     ```bash
     docker compose build
     ```

   - Run the application:

     ```bash
     docker compose up
     ```

## API Documentation

API documentation is available in the `collection/thunder-collection_Person Data API -- Go.json` file. Import this file into your API client (e.g., [Postman](https://www.postman.com/)) to explore and interact with the available endpoints.

## Testing

To run the test cases, use the following command:

```go
go test -v ./...
```
