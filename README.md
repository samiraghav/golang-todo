# GoLang Todo

GoLang Todo is a simple web application for managing todos built with GoLang and a MySQL database.

## Features

- Create new todos
- Update existing todos
- Delete todos
- Fetch all todos

## Prerequisites

Before running the application, make sure you have the following installed:

- Go programming language (version X.X.X)
- MySQL database (version X.X.X)

## Installation

1. Clone the repository:

   git clone https://github.com/samiraghav/golang-todo.git

2. Navigate to the project directory:

cd golang-todo

3. Install the dependencies:

go mod download

4. Set up the MySQL database:

Create a new database named todo_app in your MySQL server.

Modify the database connection details in the main.go file if necessary.

5. Build and run the application:

go build
./golang-todo

6. Access the application in your web browser at http://localhost:9000.

## API Endpoints
GET /todo: Fetch all todos.

POST /todo: Create a new todo.

PUT /todo/{id}: Update an existing todo.

DELETE /todo/{id}: Delete a todo.
