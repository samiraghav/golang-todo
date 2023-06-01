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

var addTodoButton = document.getElementById('add-todo-button');
addTodoButton.addEventListener('click', createTodo);

function editTodo(index) {
  var todo = todos[index];
  todoInput.value = todo.title;

  // Check if an update button already exists
  var updateButton = document.querySelector('.update-todo-button');
  if (!updateButton) {
    // Disable the add todo button
    addTodoButton.disabled = true;

    // Create a new update button if it doesn't exist
    updateButton = document.createElement('button');
    updateButton.textContent = 'Update';
    updateButton.className = 'todo-button update-todo-button';

    // Replace the add todo button with the update button
    addTodoButton.parentNode.replaceChild(updateButton, addTodoButton);

    // Disable the delete buttons of all todos
    var deleteButtons = document.querySelectorAll('.delete-todo-button');
    deleteButtons.forEach(function (button) {
      button.disabled = true;
    });

    updateButton.addEventListener('click', function () {
      var newTitle = todoInput.value.trim();
      if (newTitle !== '') {
        // Make the AJAX PUT request to update the todo
        var updateXhr = new XMLHttpRequest();
        updateXhr.open('PUT', '/todo/' + todo.id);
        updateXhr.setRequestHeader('Content-Type', 'application/json');
        updateXhr.onload = function () {
          if (updateXhr.status === 200) {
            // Todo updated successfully
            todo.title = newTitle;
            renderTodoList();
          } else {
            // Failed to update todo
            console.error('Failed to update todo');
          }
          // Enable the add todo button after updating
          addTodoButton.disabled = false;
        };

        var updatedTodo = {
          id: todo.id,
          title: newTitle,
          completed: todo.completed
        };

        updateXhr.send(JSON.stringify(updatedTodo));
      }

      todoInput.value = '';
      todoInput.parentNode.replaceChild(addTodoButton, updateButton);
      // Enable the add todo button after canceling the edit
      addTodoButton.disabled = false;
<<<<<<< HEAD

      // Enable the delete buttons of all todos
      var deleteButtons = document.querySelectorAll('.delete-todo-button');
      deleteButtons.forEach(function (button) {
        button.disabled = false;
      });

      editingIndex = -1; // Reset the editing index
=======
>>>>>>> parent of 49933b2 (Update script.js)
    });
  }
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

