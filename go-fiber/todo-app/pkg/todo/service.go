package todo

import (
	"log"
	databaseManagement "todo-app/database"
	"todo-app/entities"
)

type Service interface {
	InsertTodo(todo entities.Todo)
	UpdateTodoStatus(todoID uint, todoStatusRequest TodoStatusUpdateRequest)
	FetchTodo() []TodoResponse
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) FetchTodo() []TodoResponse {
	var todoList []entities.Todo
	err := databaseManagement.DB.Db.Find(&todoList).Error
	if err != nil {
		log.Fatal(err)
	}

	response := []TodoResponse{}

	for _, todo := range todoList {
		response = append(response, TodoResponse{
			ID:     todo.ID,
			Author: todo.Author,
			Title:  todo.Title,
		})
	}

	return response
}

func (s *service) InsertTodo(todo entities.Todo) {
	databaseManagement.DB.Db.Create(&todo)
}

func (s *service) UpdateTodoStatus(todoID uint, todoStatusRequest TodoStatusUpdateRequest) {
	databaseManagement.DB.Db.Model(&entities.Todo{}).Where("ID = ?", todoID).Update("status", todoStatusRequest.Status)
}
