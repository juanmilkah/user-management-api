# User Management API

A RESTful API built with Go (using Chi router) for user management with JWT authentication.

## Project Structure
```
go-server/
├── config/          # Configuration management
├── handlers/        # Request handlers
├── middleware/      # Custom middleware
├── models/          # Data models
├── utils/          # Utility functions
└── main.go         # Application entry point
```

## API Endpoints

### Public Endpoints

#### 1. Create User
- **Method**: POST
- **Path**: `/users`
- **Description**: Register a new user
- **Request Body**:
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
}
```
- **Validation Rules**:
  - name: required, min 2 chars, max 50 chars
  - email: required, valid email format
  - password: required, min 6 chars
- **Response**: 201 Created
```json
{
    "id": "1",
    "name": "John Doe",
    "email": "john@example.com"
}
```

#### 2. Login
- **Method**: POST
- **Path**: `/login`
- **Description**: Authenticate user and receive JWT token
- **Request Body**:
```json
{
    "email": "john@example.com",
    "password": "password123"
}
```
- **Response**: 200 OK
```json
{
    "token": "eyJhbGciOiJIUzI1..."
}
```

### Protected Endpoints
*Requires Bearer Token Authentication*

#### 1. Get All Users
- **Method**: GET
- **Path**: `/users`
- **Description**: Retrieve list of all users
- **Authentication**: Bearer token in Authorization header
- **Response**: 200 OK
```json
[
    {
        "id": "1",
        "name": "John Doe",
        "email": "john@example.com"
    },
    {
        "id": "2",
        "name": "Jane Doe",
        "email": "jane@example.com"
    }
]
```

#### 2. Get User by ID
- **Method**: GET
- **Path**: `/users/{id}`
- **Description**: Retrieve specific user by ID
- **Authentication**: Bearer token in Authorization header
- **Response**: 200 OK
```json
{
    "id": "1",
    "name": "John Doe",
    "email": "john@example.com"
}
```

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Protected endpoints require a valid JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Error Responses

The API returns consistent error responses in the following format:

```json
{
    "status": 400,
    "message": "Error message here",
    "error": "Detailed error (only in development mode)"
}
```

Common HTTP status codes:
- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

## Environment Configuration

Create a `.env` file in the root directory with the following variables:
```env
PORT=8080
JWT_SECRET=your-secret-key
ENV=development
```

## Running the Application

1. Install dependencies:
```bash
go mod tidy
```

2. Create and configure `.env` file

3. Run the server:
```bash
go run main.go
```

## Example Requests using cURL

1. Create a new user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

2. Login:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

3. Get all users (protected):
```bash
curl http://localhost:8080/users \
  -H "Authorization: Bearer <your_jwt_token>"
```

4. Get specific user (protected):
```bash
curl http://localhost:8080/users/1 \
  -H "Authorization: Bearer <your_jwt_token>"
```

## Validation Rules

### User Model
- **Name**:
  - Required
  - Minimum length: 2 characters
  - Maximum length: 50 characters
- **Email**:
  - Required
  - Must be valid email format
- **Password**:
  - Required
  - Minimum length: 6 characters

## Security Features

1. JWT Authentication
2. Password validation
3. Environment-based error responses
4. Request ID tracking
5. Panic recovery middleware
6. Logging middleware

## Dependencies

- chi (router): `github.com/go-chi/chi/v5`
- jwt-go: `github.com/golang-jwt/jwt/v5`
- validator: `github.com/go-playground/validator/v10`
- godotenv: `github.com/joho/godotenv`
```
I'll create a comprehensive README.md for your API.

```markdown
# User Management API

A RESTful API built with Go (using Chi router) for user management with JWT authentication.

## Project Structure
```
go-server/
├── config/          # Configuration management
├── handlers/        # Request handlers
├── middleware/      # Custom middleware
├── models/          # Data models
├── utils/          # Utility functions
└── main.go         # Application entry point
```

## API Endpoints

### Public Endpoints

#### 1. Create User
- **Method**: POST
- **Path**: `/users`
- **Description**: Register a new user
- **Request Body**:
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
}
```
- **Validation Rules**:
  - name: required, min 2 chars, max 50 chars
  - email: required, valid email format
  - password: required, min 6 chars
- **Response**: 201 Created
```json
{
    "id": "1",
    "name": "John Doe",
    "email": "john@example.com"
}
```

#### 2. Login
- **Method**: POST
- **Path**: `/login`
- **Description**: Authenticate user and receive JWT token
- **Request Body**:
```json
{
    "email": "john@example.com",
    "password": "password123"
}
```
- **Response**: 200 OK
```json
{
    "token": "eyJhbGciOiJIUzI1..."
}
```

### Protected Endpoints
*Requires Bearer Token Authentication*

#### 1. Get All Users
- **Method**: GET
- **Path**: `/users`
- **Description**: Retrieve list of all users
- **Authentication**: Bearer token in Authorization header
- **Response**: 200 OK
```json
[
    {
        "id": "1",
        "name": "John Doe",
        "email": "john@example.com"
    },
    {
        "id": "2",
        "name": "Jane Doe",
        "email": "jane@example.com"
    }
]
```

#### 2. Get User by ID
- **Method**: GET
- **Path**: `/users/{id}`
- **Description**: Retrieve specific user by ID
- **Authentication**: Bearer token in Authorization header
- **Response**: 200 OK
```json
{
    "id": "1",
    "name": "John Doe",
    "email": "john@example.com"
}
```

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Protected endpoints require a valid JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Error Responses

The API returns consistent error responses in the following format:

```json
{
    "status": 400,
    "message": "Error message here",
    "error": "Detailed error (only in development mode)"
}
```

Common HTTP status codes:
- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

## Environment Configuration

Create a `.env` file in the root directory with the following variables:
```env
PORT=8080
JWT_SECRET=your-secret-key
ENV=development
```

## Running the Application

1. Install dependencies:
```bash
go mod tidy
```

2. Create and configure `.env` file

3. Run the server:
```bash
go run main.go
```

## Example Requests using cURL

1. Create a new user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

2. Login:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

3. Get all users (protected):
```bash
curl http://localhost:8080/users \
  -H "Authorization: Bearer <your_jwt_token>"
```

4. Get specific user (protected):
```bash
curl http://localhost:8080/users/1 \
  -H "Authorization: Bearer <your_jwt_token>"
```

## Validation Rules

### User Model
- **Name**:
  - Required
  - Minimum length: 2 characters
  - Maximum length: 50 characters
- **Email**:
  - Required
  - Must be valid email format
- **Password**:
  - Required
  - Minimum length: 6 characters

## Security Features

1. JWT Authentication
2. Password validation
3. Environment-based error responses
4. Request ID tracking
5. Panic recovery middleware
6. Logging middleware

## Dependencies

- chi (router): `github.com/go-chi/chi/v5`
- jwt-go: `github.com/golang-jwt/jwt/v5`
- validator: `github.com/go-playground/validator/v10`
- godotenv: `github.com/joho/godotenv`
```

