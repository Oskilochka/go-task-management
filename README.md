# Task Management System

## Overview

The **Task Management System** is a RESTful API built using the Go language.
This system allows users to perform CRUD operations on tasks, making it ideal for managing to-do lists or tracking
project tasks.
The project demonstrates the use of several Go libraries and approaches to build a robust, modular, and maintainable web
service.

## Features
- **User Authentication:**
    - Register a new user with username, email, and password.
    - Log in to receive a JSON Web Token (JWT) for authenticated requests.
- **JWT-Based Authentication:** Secure endpoints with JWT, allowing access only to authenticated users.
- **Create Tasks:** Add new tasks with details such as title, description, status of completion.
- **Read Tasks:** Retrieve a list of all tasks or specific tasks by ID.
- **Update Tasks:** Modify existing task details.
- **Delete Tasks:** Remove tasks from the system.
- **User-specific Tasks:** Each user can only view, create, update, or delete their own tasks.

## Project Structure

### Description of Key Components

- **`main.go`:** The entry point of the application that initializes the database and sets up the HTTP server and
  routes.
- **`/handlers`:** Contains HTTP handler functions that define the logic for processing API requests.
- **`/models`:** Defines data models that map to the database tables using the GORM library.
- **`/routes`:** Configures and initializes the API routes and associates them with the respective handlers.
- **`/database`:** Manages database connections and performs automatic migrations.
- **`/utils`:** Provides utility functions, such as sending standardized JSON responses.
- **`/auth`:** Handles JWT token generation and verification logic.
- **`/middlewares`:** Provides middleware functions for handling JWT authentication.


## Libraries and Approaches Used

### Libraries

1. **`gorilla/mux`**
    - A powerful HTTP router and URL matcher for building Go web servers.
    - **Usage:** This library is used to handle HTTP request routing, allowing the definition of RESTful endpoints with
      ease and flexibility.
    - **Reason:** Its robust routing capabilities, including dynamic URL parameters and method-based routing, make it an
      ideal choice for building REST APIs.

2. **`gorm.io/gorm`**
    - An Object-Relational Mapping (ORM) library for Go.
    - **Usage:** It simplifies database interactions by allowing developers to work with Go structs instead of raw SQL
      queries.
    - **Reason:** `GORM` provides a high-level API for database operations, including migrations, which greatly reduces
      the boilerplate code needed for CRUD operations.

3. **`gorm.io/driver/sqlite`**
    - A GORM driver for the SQLite database.
    - **Usage:** This library enables the use of SQLite as the database engine for the application.
    - **Reason:** SQLite is lightweight, serverless, and easy to set up, making it ideal for small projects, prototypes,
      and local development.

4. **`golang-jwt/jwt/v5`**
    - A Go library for working with JSON Web Tokens (JWT).
    - **Usage:** This library is used to generate and verify JWTs for securing user authentication.
    - **Reason:** JWT allows secure transmission of information between the client and server, ensuring that only authenticated users can access certain endpoints.

5. **`bcrypt` (`golang.org/x/crypto/bcrypt`)**
    - A password hashing library.
    - **Usage:** It is used to securely hash and check user passwords during registration and login.
    - **Reason:** Bcrypt is known for its strong security against brute-force attacks and is widely used for password management.


### Approaches

- **Modular Structure:** The project is organized into different packages (`handlers`, `models`, `routes`, `database`,
  `utils`) to separate concerns and improve maintainability.
- **RESTful API Design:** The API follows REST principles, providing clear and standardized endpoints for managing
  tasks.
- **JSON-based Communication:** All data exchange between the client and server is done using JSON, which is a widely
  used format for APIs due to its simplicity and readability.
- **Error Handling:** The project includes consistent error handling throughout the application, ensuring that
  meaningful error messages and appropriate HTTP status codes are returned.
- **Automated Migrations:** Using `GORM`'s auto-migration feature, the database schema is automatically updated based on
  the defined models, simplifying schema management.
- **JWT Authentication:** Protects routes using JWT middleware, ensuring that only authenticated users can create, read, update, or delete tasks.

## Getting Started

### Prerequisites

- **Go:** Make sure Go is installed on your system. You can download it
  from [the official website](https://golang.org/dl/).

### Installation
1. **Clone the Repository:**

   ```bash
   git clone https://github.com/Oskilochka/go-task-management.git
   cd task-management-system
   ```

2. **Install Dependencies:**

    The project uses Go modules. Simply run:
    
   ```bash
   go mod tidy
    ```

3. **Run the Application:**

    Start the server by running:
   ```bash
   go run main.go
    ```
The server will start on http://localhost:8080.

### Docker Support
This project includes a Dockerfile and a docker-compose.yml file,
allowing you to easily build and run the application in a containerized environment.

#### Running the Application with Docker
1. Build the Docker image and start the container:
     ```bash
       docker-compose up --build
    ```
2. Access the API: The application will be accessible at http://localhost:8080.

***Notes***: Ensure that Docker is installed and running on your machine before executing the commands.

### API Endpoints

**GET /tasks** - Retrieve all tasks.

**GET /tasks/{id}** - Retrieve a task by ID.

**POST /tasks** - Create a new task.

**PUT /tasks/{id}** - Update an existing task.

**DELETE /tasks/{id}** - Delete a task by ID.

## API Endpoints
To interact with the API, you can use tools like curl, Postman, or any HTTP client library.

### Authentication Endpoints

#### Register a new user

- **Endpoint**: `POST /register`
- **Description**: Allows a new user to register by providing a username, email, and password.
- **Request Body**:
    ```json
    {
        "username": "exampleuser",
        "email": "example@example.com",
        "password": "password123"
    }
    ```
- **Response**: A success message upon successful registration.

#### User Login

- **Endpoint**: `POST /login`
- **Description**: Logs in a user by generating a JWT token.
- **Request Body**:

    ```json
    {
        "username": "exampleuser",
        "password": "password123"
    }
    ```

- **Response**: A JWT token to be used for authenticated requests.

    ```json
    {
        "token": "<JWT_TOKEN>"
    }
    ```

### Tasks API Examples

1. **Create a Task:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -d '{
        "title": "My New Task Title",
        "description": "Description of the new task",
        "completed": false
      }'
   ```

2.**Get all Tasks:**
```bash
curl http://localhost:8080/tasks \
 -H "Authorization: Bearer <JWT_TOKEN>" 
   ```

3.**Get a Task by ID:**

Replace {id} with the actual task ID.
```bash
curl -X GET http://localhost:8080/tasks/{id} \
-H "Authorization: Bearer <JWT_TOKEN>" 
```

4.**Update a Task by ID:**

Replace {id} with the ID of the task you want to update. 

For example, if the ID is 1, the command would be:
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -d '{
        "title": "Updated Task Title",
        "description": "Updated description of the task",
        "completed": true
      }'
 ```

5.**Delete a Task by ID:**
   
Replace {id} with the actual task ID.

```bash
curl -X DELETE http://localhost:8080/tasks/{id} \
  -H "Authorization: Bearer <JWT_TOKEN>" \
   ```

6.**Register:**

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
        "username": "username",
        "password": "your password",
        "email": "test@gmail.com"
      }'  
   ```
 
7. **Login:**
```bash
curl -X POST http://localhost:8080/login \
    -H "Content-Type: application/json" \
    -d '{
        "username": "username",
        "password": "your password"
        }'
   ```
