# GoLang Todo

GoLang Todo is a simple web application for managing todos built with GoLang and a MySQL database.

## Features

- Create new todos
- Update existing todos
- Delete todos
- Fetch all todos

## Prerequisites

Before running the application, make sure you have the following installed:

- Go programming language 
- MySQL database

## Installation

1. Clone the repository:
  
   ```
   git clone https://github.com/samiraghav/golang-todo.git
   ```

2. Navigate to the project directory:

   ```
   cd golang-todo
   ```

3. Navigate to the backend folder in golang-todo directory

   ```
   cd backend
   ```

3. Install the dependencies:

   ```
   go mod download
   ```

4. Set up the MySQL database:

   Create a new database named todo_app in your MySQL server.

   Modify the database connection details in the main.go file if necessary.

5. Run the application:

   ```
   go run main.go
   ```

6. Access the application in your web browser at 

   ```
   http://localhost:9000
   ```

## API Endpoints
- GET /todo: Fetch all todos.

- POST /todo: Create a new todo.

- PUT /todo/{id}: Update an existing todo.

- DELETE /todo/{id}: Delete a todo.
