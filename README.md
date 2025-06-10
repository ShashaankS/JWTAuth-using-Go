# JWTAuth-using-Go 🚀

A practical, modular example of JWT (JSON Web Token) authentication in Go, using the Gin web framework, GORM ORM, Bcrypt for password hashing, and JWT for secure authentication. This project is ideal for learning or as a starting point for your own secure Go APIs! 🔐

---

## 📚 Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [How JWT Authentication Works](#how-jwt-authentication-works)
- [Example Middleware](#example-middleware)
- [Environment Variables](#environment-variables)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

---

## ✨ Features

- User registration and login with JWT token generation
- Passwords securely hashed with Bcrypt
- Protected routes accessible only with valid JWT tokens
- Modular project structure (controllers, middleware, models)
- Environment-based configuration

---

## 🛠️ Tech Stack

| Library  | Purpose                       |
|----------|------------------------------|
| Gin      | Web framework & routing 🚦   |
| GORM     | ORM for database access 🗄️   |
| Bcrypt   | Password hashing 🔑          |
| JWT      | Token-based authentication 🪪 |

---

## 🗂️ Project Structure
```
JWTAuth-using-Go/
├── controllers/ # Request handlers (auth logic)
├── initializers/ # App initialization (DB, env)
├── middleware/ # JWT authentication middleware
├── models/ # Data models (User struct, etc.)
├── main.go # Entry point
├── .env.example # Example environment variables
├── go.mod # Go module definition
├── go.sum # Go module checksums
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18+ 🦫
- (Optional) Database, if you want persistent storage (I used Neon DB)

### Installation

1. **Clone the repo**
```
git clone https://github.com/ShashaankS/JWTAuth-using-Go.git
cd JWTAuth-using-Go
```

2. **Copy and edit environment variables**
```
cp .env.example .env
```

3. **Install dependencies**
```
go mod tidy
```

4. **Run the app**
```
go run main.go
```

---

## 🧪 Usage

- **Register:** `POST /register` with user credentials
- **Login:** `POST /login` with valid credentials to receive a JWT
- **Protected Routes:** Include your JWT in the `Authorization: Bearer <token>` header to access secured endpoints

---

## 🔑 How JWT Authentication Works

1. **User registers:** Password is hashed with Bcrypt and stored securely in the database via GORM.
2. **User logs in:** On successful login, the server issues a JWT containing user claims.
3. **Client stores JWT:** Usually in localStorage or a cookie.
4. **Accessing protected routes:** The client sends the JWT in the `Authorization` header.
5. **Server validates JWT:** Middleware checks the token before granting access.

---
## ⚙️ Environment Variables

Edit your `.env` using `.env.example` as a template:

- `SECRET` — Secret key for signing JWTs 🔑
- Database connection strings (if needed)

---

## 📦 Dependencies

- [Gin](https://gin-gonic.com/) — Web framework
- [GORM](https://gorm.io/) — ORM for Go
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — Password hashing
- [JWT-Go](https://github.com/golang-jwt/jwt) — JWT library for Go

---


## 🙏 Acknowledgements

- Inspired by community best practices for JWT in Go
- Uses patterns from the Go ecosystem

---

> **Note:** This project is for educational use and may need enhancements for production (better error handling, token refresh, secure storage, etc.) 😎

---
