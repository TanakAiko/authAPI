# 🔐 AuthAPI

<div align="center">

A lightweight, secure authentication API built with Go that provides user registration, login, session management, and user data retrieval capabilities.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=for-the-badge&logo=sqlite&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![bcrypt](https://img.shields.io/badge/Security-bcrypt-green?style=for-the-badge)
![UUID](https://img.shields.io/badge/Sessions-UUID-orange?style=for-the-badge)

</div>

<details>
<summary>📋 Table of Contents</summary>

- [🔐 AuthAPI](#-authapi)
  - [✨ Features](#-features)
  - [🔧 Prerequisites](#-prerequisites)
  - [📦 Installation](#-installation)
    - [Option 1: Local Development](#option-1-local-development)
    - [Option 2: Docker](#option-2-docker)
  - [🚀 Quick Start](#-quick-start)
  - [📚 API Reference](#-api-reference)
    - [1. Register](#1-register)
    - [2. Login](#2-login)
    - [3. Authorized](#3-authorized)
    - [4. Logout](#4-logout)
    - [5. Get User Data](#5-get-user-data)
  - [📁 Project Structure](#-project-structure)
  - [🛠️ Development](#️-development)
    - [Dependencies](#dependencies)
    - [Building from Source](#building-from-source)
    - [Running with Docker](#running-with-docker)
  - [🧪 Testing](#-testing)
    - [1. Register a New User](#1-register-a-new-user)
    - [2. Login](#2-login-1)
    - [3. Check Authorization (Use sessionID from login response)](#3-check-authorization-use-sessionid-from-login-response)
    - [4. Get User Data](#4-get-user-data)
    - [5. Logout](#5-logout)
  - [🤝 Contributing](#-contributing)

</details>

## ✨ Features

- **User Registration**: Secure user account creation with password hashing (bcrypt)
- **User Authentication**: Login with email/nickname and password
- **Session Management**: UUID-based sessions with 24-hour expiration
- **Authorization Validation**: Verify active sessions
- **User Data Retrieval**: Get authenticated user information
- **Secure Logout**: Session termination and cleanup
- **SQLite Database**: Lightweight, embedded database storage
- **Docker Support**: Containerized deployment ready

## 🔧 Prerequisites

- Go 1.20 or higher
- SQLite3
- Docker (optional, for containerized deployment)

## 📦 Installation

### Option 1: Local Development

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd authAPI
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

   Or install packages individually:
   ```bash
   go get golang.org/x/crypto/bcrypt
   go get github.com/mattn/go-sqlite3
   go get github.com/google/uuid
   ```

3. **Run the server**
   ```bash
   go run main.go
   ```
   
   > **Note**: The database will be automatically created on first run in `./databases/auth.db`

### Option 2: Docker

1. **Build the Docker image**
   ```bash
   docker build -t authapi .
   ```

2. **Run the container**
   ```bash
   docker run -p 8081:8081 -v $(pwd)/databases:/app/databases authapi
   ```

## 🚀 Quick Start

1. **Start the server**
   ```bash
   go run main.go
   ```

2. **Server will be available at**
   ```
   http://localhost:8081
   ```

3. **Test the API**
   ```bash
   # Register a new user
   curl -X POST http://localhost:8081/ \
     -H "Content-Type: application/json" \
     -d '{
       "action": "register",
       "body": {
         "nickname": "johndoe",
         "age": 30,
         "gender": "male",
         "firstName": "John",
         "lastName": "Doe",
         "email": "john.doe@example.com",
         "password": "securePassword123"
       }
     }'
   ```

## 📚 API Reference

All requests are sent as POST to the root endpoint (`/`) with a JSON body containing an `action` field and a `body` field.

### 1. Register

Create a new user account.

**Request:**
```json
{
  "action": "register",
  "body": {
    "nickname": "string",
    "age": "int",
    "gender": "string",
    "firstName": "string",
    "lastName": "string",
    "email": "string",
    "password": "string"
  }
}
```

**Response:**
- **Status**: `201 Created`
- **Body**: `"New user created"`

**Example:**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "register",
    "body": {
      "nickname": "exampleNickname",
      "age": 30,
      "gender": "male",
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "password": "securePassword123"
    }
  }'
```

---

### 2. Login

Authenticate a user and receive session information.

**Request:**
```json
{
  "action": "login",
  "body": {
    "identifier": "string (email or nickname)",
    "password": "string"
  }
}
```

**Response:**
- **Status**: `200 OK`
- **Body**: User data object with session information

**Example:**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "login",
    "body": {
      "identifier": "john.doe@example.com",
      "password": "securePassword123"
    }
  }'
```

---

### 3. Authorized

Verify if a session is valid.

**Request:**
```json
{
  "action": "authorized",
  "body": {
    "sessionID": "string (UUID)"
  }
}
```

**Response:**
- **Status**: `202 Accepted`
- **Body**: `"The session is valid"`

**Example:**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "authorized",
    "body": {
      "sessionID": "6a09a3da-26ee-4b35-870c-d7a4f22f939c"
    }
  }'
```

---

### 4. Logout

Terminate a user session.

**Request:**
```json
{
  "action": "logout",
  "body": {
    "sessionID": "string (UUID)"
  }
}
```

**Response:**
- **Status**: `200 OK`
- **Body**: `"The session is deleted"`

**Example:**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "logout",
    "body": {
      "sessionID": "6a09a3da-26ee-4b35-870c-d7a4f22f939c"
    }
  }'
```

---

### 5. Get User Data

Retrieve authenticated user information.

**Request:**
```json
{
  "action": "getUserData",
  "body": {
    "sessionID": "string (UUID)"
  }
}
```

**Response:**
- **Status**: `200 OK`
- **Body**: User data object

**Example:**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "getUserData",
    "body": {
      "sessionID": "6a09a3da-26ee-4b35-870c-d7a4f22f939c"
    }
  }'
```

## 📁 Project Structure

```
authAPI/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies checksum
├── Dockerfile              # Docker configuration
├── README.md               # This file
├── databases/              # Database files and SQL scripts
│   └── sqlRequests/
│       ├── createTable.sql
│       ├── insertNewSession.sql
│       └── insertNewUser.sql
├── internals/              # Internal application logic
│   ├── dbManager/          # Database initialization
│   │   └── initDB.go
│   ├── handlers/           # HTTP request handlers
│   │   ├── mainHandler.go
│   │   ├── registerHandler.go
│   │   ├── loginHandler.go
│   │   ├── authorized.go
│   │   ├── logoutHandler.go
│   │   └── getUserDataHandler.go
│   └── tools/              # Utility functions
│       └── utils.go
├── models/                 # Data models
│   ├── user.go
│   ├── session.go
│   └── request.go
└── script/                 # Utility scripts
    ├── init.sh
    └── push.sh
```

## 🛠️ Development

### Dependencies

This project uses the following Go packages:

- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)**: Password hashing and verification
- **[go-sqlite3](https://github.com/mattn/go-sqlite3)**: SQLite database driver
- **[uuid](https://github.com/google/uuid)**: UUID generation for sessions

### Building from Source

```bash
# Build the binary
go build -o authapi-server

# Run the binary
./authapi-server
```

### Running with Docker

```bash
# Build and run with Docker
docker build -t authapi .
docker run -p 8081:8081 -v $(pwd)/databases:/app/databases authapi
```

## 🧪 Testing

Complete test suite with example commands:

### 1. Register a New User
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "register",
    "body": {
      "nickname": "testuser",
      "age": 25,
      "gender": "female",
      "firstName": "Jane",
      "lastName": "Smith",
      "email": "jane.smith@example.com",
      "password": "testPassword456"
    }
  }'
```

### 2. Login
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "login",
    "body": {
      "identifier": "jane.smith@example.com",
      "password": "testPassword456"
    }
  }'
```

### 3. Check Authorization (Use sessionID from login response)
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "authorized",
    "body": {
      "sessionID": "YOUR_SESSION_ID_HERE"
    }
  }'
```

### 4. Get User Data
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "getUserData",
    "body": {
      "sessionID": "YOUR_SESSION_ID_HERE"
    }
  }'
```

### 5. Logout
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "logout",
    "body": {
      "sessionID": "YOUR_SESSION_ID_HERE"
    }
  }'
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

**Note**: This API is designed for educational purposes. For production use, consider additional security measures such as HTTPS, rate limiting, input validation, and comprehensive error handling.

---

<div align="center">

**⭐ Star this repository if you found it helpful! ⭐**

Made with ❤️ from 🇸🇳

</div>