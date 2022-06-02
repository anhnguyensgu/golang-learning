package main

import (
	"fmt"
	"todo-app/api/routes"
	database_management "todo-app/database"
	"todo-app/pkg/todo"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello world")
	database_management.SetupDatabase()

	app := fiber.New()

	s := todo.NewService()
	routes.TodoRoute(app, s)
	app.Listen(":3010")
}
