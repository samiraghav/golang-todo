var todoList = document.getElementById('list');
var todoInput = document.getElementById('input');
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
  todoInput.value = todo.title;
  todos.splice(index, 1);
  renderTodoList();
}

function deleteTodo(index) {
  todos.splice(index, 1);
  renderTodoList();
}