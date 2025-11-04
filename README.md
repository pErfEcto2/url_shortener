# URL Shortener

A lightweight URL shortener web application built with Go, Gin, and HTMX with JWT authentication.

## Features

- User authentication with JWT tokens
- URL shortening with unique short codes
- User dashboard for managing shortened URLs
- HTMX for dynamic frontend interactions
- In-memory data storage

## Tech Stack

- Backend: Go, Gin web framework
- Frontend: HTML, HTMX
- Authentication: JWT
- Styling: Tacit CSS
- Storage: In-memory (extensible to databases)

## Project Structure

```
url_shortener/
├── cmd/main.go                 # Application entry point
├── internal/
│   ├── auth/                   # JWT authentication
│   ├── db/                     # Data storage
│   ├── handlers/               # HTTP handlers
│   ├── models/                 # Data models
│   └── shortener/              # URL shortening logic
└── static/                     # Frontend assets
    ├── css/                    # Stylesheets
    ├── js/                     # JavaScript
    └── *.html                  # HTML templates
```

##  API Endpoints

- GET / - Home page
- GET /signup - Registration page
- POST /signup - Create account
- GET /login - Login page
- POST /login - Authenticate user
- GET /user - User dashboard (protected)
- POST /shorten - Create shortened URL

## Setup

Create .env file:
```
HOST=localhost
PORT=8080
SECRET=your-jwt-secret-key
```

Install dependencies:

```
go mod tidy
```

Run the application:

```
go run ./cmd
```

Visit http://localhost:8080

## Usage

- Public: Shorten URLs without account on home page
- Registered Users: Sign up, log in, and manage shortened URLs in dashboard

## TODO

- signup and login input validation and sanitization
- signup and login rate limits
- url validation before shortening
- validate signup password complexity
- add validation that shortened URLs are unique across all users
