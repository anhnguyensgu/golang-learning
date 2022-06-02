package handlers

import (
	"log"
	"todo-app/entities"
	"todo-app/pkg/todo"

	"github.com/gofiber/fiber/v2"
)

func FetchTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(service.FetchTodo())
	}
}

func InsertBook(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todo := entities.Todo{}

		err := c.BodyParser(&todo)
		if err != nil {
			log.Fatal(err)
		}

		service.InsertTodo(todo)
		return c.SendString("To be continue")
	}
}
