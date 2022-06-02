package routes

import (
	"todo-app/api/handlers"
	"todo-app/pkg/todo"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(app *fiber.App, service todo.Service) {
	todoPath := "/todo"
	app.Get(todoPath, handlers.FetchTodo(service))
	app.Post(todoPath, handlers.InsertBook(service))
}
