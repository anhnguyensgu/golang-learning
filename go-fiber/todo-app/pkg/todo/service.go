package todo

import (
	"log"
	databaseManagement "todo-app/database"
	"todo-app/entities"
)

type Service interface {
	InsertTodo(todo entities.Todo)
	FetchTodo() []entities.Todo
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) FetchTodo() []entities.Todo {
	var todoList []entities.Todo
	err := databaseManagement.DB.Db.Find(&todoList).Error
	if err != nil {
		log.Fatal(err)
	}
	return todoList
}

func (s *service) InsertTodo(todo entities.Todo) {
	databaseManagement.DB.Db.Create(&todo)
}
