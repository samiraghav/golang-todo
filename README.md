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

   git clone https://github.com/samiraghav/golang-todo.git
   <!-- Clone the repository to your local machine -->

2. Navigate to the project directory:

   cd golang-todo
   <!-- Change your working directory to the project directory -->

3. Install the dependencies:

   go mod download
   <!-- Download and install the required dependencies -->

4. Set up the MySQL database:

   Create a new database named todo_app in your MySQL server.

   Modify the database connection details in the main.go file if necessary.
   <!-- Create a new database and configure the connection details in `main.go` if required -->

5. Build and run the application:

   ./golang-todo
   <!-- Execute the application -->

6. Access the application in your web browser at http://localhost:9000.
<!-- Open your web browser and visit the provided URL to access the application -->

## API Endpoints
- GET /todo: Fetch all todos.

- POST /todo: Create a new todo.

- PUT /todo/{id}: Update an existing todo.

- DELETE /todo/{id}: Delete a todo.
