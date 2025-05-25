# 🏡 Domestic Help Management Backend

![Go Version](https://img.shields.io/badge/Go-1.21-blue)  
A secure and efficient backend built with the **Gin framework** to manage domestic help services. It provides a RESTful API for handling users, services, and operations with **JWT authentication** and **PostgreSQL** integration.  

---

## 🌟 Features

- 🛡️ **Secure Authentication**: JWT-based middleware to protect routes.  
- ⚡ **Fast and Scalable**: Built using the Gin web framework for high performance.  
- 💾 **Database Integration**: PostgreSQL for data storage with **GORM** for ORM.  
- ☁️ **Cloud Support**: Ready for AWS S3 integration for file storage.  
- 📄 **Auto Migration**: Automatically syncs database schema with models.  

---

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+**: [Install Go](https://golang.org/doc/install)  
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/)  

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/2deadmen/domestic_backend.git
   cd domestic_backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure your environment variables in a `.env` file:
   ```env
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASS=12345
   DB_NAME=domestic
   DB_PORT=5432
   JWT_SECRET=your-secret-key
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

---

## 📂 Project Structure

```
domestic_backend/
├── main.go                 # Application entry point
├── models/                 # Database models
│   └── user.go             # User model and database operations
├── services/               # Database and other service configurations
│   └── db.go               # PostgreSQL connection and initialization
├── middlewares/            # Custom middleware
│   └── jwt.go              # JWT authentication middleware
├── routes/                 # API route definitions
│   └── employer_routes.go  # Employer route handling
│   └── employee_routes.go  # Employee route handling
│   └── jobcard_routes.go   # Job card route handling
│   └── jobapp_routes.go    # Job application route handling
├── utils/                  # Utility functions
│   └── jwt.go              # JWT generation and helpers
├── docs/                   # Swagger documentation setup
└── go.mod                  # Go module dependencies
```

---

## 🔑 Authentication

This project uses JWT tokens for secure authentication.  

### Generating a Token
Use the `/login` endpoint (to be implemented) to get a JWT token.  

### Securing Routes
All routes are protected using the JWT middleware. Example:
```bash
curl -H "Authorization: Bearer <your-jwt-token>" http://localhost:8080/users
```

---

## 🛠️ API Endpoints

### Employer Routes
Group: `/employers`

| Method | Endpoint               | Description                                |
|--------|------------------------|--------------------------------------------|
| POST   | `/`                    | Register a new employer.                  |
| POST   | `/verify-otp`          | Verify OTP for employer account.          |
| POST   | `/sign-in`             | Authenticate employer and return JWT.     |
| GET    | `/`                    | Get a list of all employers.              |
| GET    | `/{id}`                | Retrieve details of an employer by ID.    |
| PUT    | `/{id}`                | Update employer details by ID.            |
| DELETE | `/{id}`                | Delete employer by ID.                    |

### Employee Routes
Group: `/employees`

| Method | Endpoint               | Description                                |
|--------|------------------------|--------------------------------------------|
| POST   | `/`                    | Register a new employee.                  |
| POST   | `/sign-in`             | Authenticate employee and return JWT.     |
| GET    | `/`                    | Get a list of all employees.              |
| GET    | `/{id}`                | Retrieve details of an employee by ID.    |
| PUT    | `/{id}`                | Update employee details by ID.            |
| DELETE | `/{id}`                | Delete employee by ID.                    |

### Job Card Routes
Group: `/jobcards`

| Method | Endpoint               | Description                                |
|--------|------------------------|--------------------------------------------|
| POST   | `/`                    | Create a new job card.                    |
| GET    | `/`                    | Get a list of all job cards.              |
| GET    | `/{id}`                | Retrieve details of a job card by ID.     |
| PUT    | `/{id}`                | Update job card details by ID.            |
| DELETE | `/{id}`                | Delete job card by ID.                    |
| GET    | `/active`              | Get all active job cards.                 |
| PUT    | `/{id}/active`         | Update the active status of a job card.   |

### Job Application Routes
Group: `/job-applications`

| Method | Endpoint               | Description                                |
|--------|------------------------|--------------------------------------------|
| POST   | `/`                    | Create a new job application.             |
| DELETE | `/{id}`                | Delete a job application by ID.           |
| PUT    | `/applications/{id}/status` | Update the status of a job application. |

---

## 🧪 Running Tests

To run tests (if implemented):
```bash
go test ./...
```

---

## 🌐 Deployment

1. Build the application:
   ```bash
   go build -o domestic_backend
   ```

2. Deploy the binary to your server or containerize it with Docker. Example Dockerfile:
   ```dockerfile
   FROM golang:1.21
   WORKDIR /app
   COPY . .
   RUN go mod tidy
   RUN go build -o domestic_backend
   CMD ["./domestic_backend"]
   ```

---

## 🎯 Future Enhancements

- 🔄 Add support for task scheduling and notifications.  
- 📤 Integrate AWS S3 for file uploads.  
- 📈 Implement metrics and logging.  
- 🌍 Internationalization (i18n).  

---

## 🤝 Contributing

Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request.  

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.  

---

## ❤️ Acknowledgements

- [Gin Web Framework](https://gin-gonic.com/)  
- [GORM ORM](https://gorm.io/)  
- [PostgreSQL](https://www.postgresql.org/)  
- [JWT for Go](https://github.com/golang-jwt/jwt)  

---

## 📝 Author

Developed with ❤️ by [2deadmen](https://github.com/2deadmen).

