# 🏥 Hospital Management System API (Go + Gin)

A REST API built in Golang to manage patients in a hospital. Uses Gin, GORM, PostgreSQL, and JWT for authentication.

---

## 🚀 Features

- JWT login (`/login`)
- Role-based access (doctor & receptionist)
- Secure CRUD APIs for patient management
- PostgreSQL with GORM ORM
- Middleware protection for routes

---

 Demo Credentials

 Doctor
```json
{
  "username": "doctor1",
  "password": "pass123"
}

Receptionist
{
  "username": "reception1",
  "password": "pass123"
}

Method	Route	Access
POST	/login	      Public
POST	/patients	    Doctor, Receptionist
GET	    /patients	    Doctor, Receptionist
GET	    /patients/:id	Doctor, Receptionist
PUT	    /patients/:id	Doctor only
DELETE	/patients/:id	Doctor only

Setup Instructions
Clone the repo

Create a .env file in root folder:
DB_URL=your_postgres_url
JWT_SECRET=secretkey


Install Go dependencies:
go mod tidy


Run the server:
go run main.go

Test routes using Postman. Use the JWT token from /login in the Authorization header (type: Bearer Token).

📦 Tech Stack
Go (Golang)
Gin framework
PostgreSQL (Railway)
GORM
JWT Authentication