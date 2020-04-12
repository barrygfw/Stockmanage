package todolist

import "graduationProjectPeng/models/todolistModel"

func AddToDo(todo *todolistModel.Todolist) error {
	return todo.AddToDo()
}
