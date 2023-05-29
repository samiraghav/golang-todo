var todoList = document.getElementById('list');
var todoInput = document.getElementById('new-todo-input');
var todos = [];

function renderTodoList() {
  todoList.innerHTML = '';
  todos.forEach(function (todo, index) {
    var li = document.createElement('li');
    li.className = 'todo-item ' + (todo.completed ? 'completed' : '');
    li.innerHTML = '<span class="todo-item-text ' + (todo.completed ? 'completed' : '') + '">' + todo.title + '</span>' +
      '<div class="todo-item-actions">' +
      '<button type="button" class="edit-todo-button" onclick="editTodo(' + index + ')">Edit</button>' +
      '<button type="button" class="delete-todo-button" onclick="deleteTodo(' + index + ')">Delete</button>' +
      '</div>';
    todoList.appendChild(li);
  });
}

function addTodo() {
  var title = todoInput.value.trim();
  if (title === '') {
    todoInput.classList.add('error');
  } else {
    todoInput.classList.remove('error');
    var todo = {
      title: title,
      completed: false
    };
    todos.push(todo);
    renderTodoList();
    todoInput.value = '';
  }
}

function editTodo(index) {
  var todo = todos[index];
  var newTitle = todoInput.value.trim();

  // Make sure the new title is not empty
  if (newTitle === '') {
    todoInput.classList.add('error');
    return;
  }

  // Make the AJAX PUT request
  var xhr = new XMLHttpRequest();
  xhr.open('PUT', '/todo/' + todo.id);
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.onload = function () {
    if (xhr.status === 200) {
      // Todo updated successfully
      todo.title = newTitle;
      renderTodoList();
      todoInput.value = '';
    } else {
      // Failed to update todo
      console.error('Failed to update todo');
    }
  };
  xhr.send(JSON.stringify({ title: newTitle, completed: todo.completed }));
}


function deleteTodo(index) {
  var todo = todos[index];

  // Make the AJAX DELETE request
  var xhr = new XMLHttpRequest();
  xhr.open('DELETE', '/todo/' + todo.id);
  xhr.onload = function () {
    if (xhr.status === 200) {
      // Todo deleted successfully
      todos.splice(index, 1);
      renderTodoList();
    } else {
      // Failed to delete todo
      console.error('Failed to delete todo');
    }
  };
  xhr.send();
}


// Function to make an AJAX POST request to create a new todo
function createTodo() {
  // Get the todo title from the input field
  var newTodoInput = document.getElementById('new-todo-input');
  var title = newTodoInput.value.trim();

  // Send the AJAX request only if the title is not empty
  if (title !== '') {
    // Create a todo object with the title
    var todo = {
      title: title,
      completed: false
    };

    // Make the AJAX POST request
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/todo');
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onload = function () {
      if (xhr.status === 201) {
        // Todo created successfully
        var response = JSON.parse(xhr.responseText);
        var todoId = response.todo_id;

        // Add the new todo to the frontend list
        var newTodo = {
          id: todoId,
          title: title,
          completed: false
        };
        todos.push(newTodo);
        renderTodoList();

        // Clear the input field
        newTodoInput.value = '';
      } else {
        // Failed to create todo
        console.error('Failed to create todo');
      }
    };
    xhr.send(JSON.stringify(todo));
  }
}

// Function to make an AJAX GET request to fetch all todos
function fetchTodos() {
  // Make the AJAX GET request
  var xhr = new XMLHttpRequest();
  xhr.open('GET', '/todo');
  xhr.onload = function () {
    if (xhr.status === 200) {
      // Todos fetched successfully
      var response = JSON.parse(xhr.responseText);
      todos = response.data;
      renderTodoList();
    } else {
      // Failed to fetch todos
      console.error('Failed to fetch todos');
    }
  };
  xhr.send();
}

// Call fetchTodos() when the page is loaded to retrieve todos
fetchTodos();


// Event listener for the add todo button click
var addTodoButton = document.getElementById('add-todo-button');
addTodoButton.addEventListener('click', createTodo);
