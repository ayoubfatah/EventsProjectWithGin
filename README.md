# Go Gin REST API

A simple REST API built with **Go** and the **Gin Web Framework**.

This project is a small backend application created to practice building RESTful APIs with Go, including routing, HTTP requests, JSON responses, and CRUD operations.

## 🚀 Features

* RESTful API endpoints
* Built with Go
* Gin web framework
* JSON request and response handling
* CRUD operations
* Simple and clean project structure

## 🛠️ Technologies

* **Go**
* **Gin**


## ⚙️ Getting Started

### Prerequisites

Make sure you have **Go** installed on your machine.

Check your Go version:

```bash
go version
```

### Clone the repository

```bash
git clone <your-repository-url>
cd <project-directory>
```

### Install dependencies

```bash
go mod download
```

### Run the API

```bash
go run main.go
```

The server should start on:

```text
http://localhost:8080
```

## 📡 API Endpoints

### Events

| Method | Endpoint      | Description        |
| ------ | ------------- | ------------------ |
| GET    | `/events`     | Get all events     |
| GET    | `/events/:id` | Get an event by ID |

### Registration

| Method | Endpoint            | Description               |
| ------ | ------------------- | ------------------------- |
| PUT    | `/registration/:id` | Register for an event     |
| DELETE | `/registration/:id` | Cancel event registration |

### Users

| Method | Endpoint  | Description       |
| ------ | --------- | ----------------- |
| POST   | `/signup` | Create a new user |
| GET    | `/users`  | Get all users     |
| POST   | `/login`  | Log in a user     |

> **Note:** The create, update, and delete event routes are currently commented out in `RegisterRoutes`.

## 🧪 Testing the API

You can test the API using tools such as Postman, Insomnia, or `curl`.

### Get all events

```bash
curl http://localhost:8080/events
```

### Get an event by ID

```bash
curl http://localhost:8080/events/1
```

### Create a user

```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### Get all users

```bash
curl http://localhost:8080/users
```

### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### Register for an event

```bash
curl -X PUT http://localhost:8080/registration/1
```

### Cancel registration

```bash
curl -X DELETE http://localhost:8080/registration/1
```


## 🧪 Testing the API

You can test the endpoints using tools such as:

* Postman
* Insomnia
* cURL
* HTTPie

Example:

```bash
curl http://localhost:8080/...
```

## 📌 Purpose

This project was built as a simple introduction to creating REST APIs with Go and Gin. It can also serve as a starting point for building more complete Go backend applications.

## 📄 License

This project is open source and available under the MIT License.
