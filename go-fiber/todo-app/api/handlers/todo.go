package handlers

import (
	"fmt"
	"log"
	"strconv"
	"todo-app/entities"
	"todo-app/pkg/todo"

	"github.com/gofiber/fiber/v2"
)

func FetchTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(service.FetchTodo())
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

func UpdateTodoStatus(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todoIDstr := c.Params("todoID")
		log.Printf("Updating status %s", todoIDstr)
		todoID, err := strconv.Atoi(todoIDstr)
		if err == nil {
			fmt.Printf("%q looks like a number.\n", todoIDstr)
		}

		todoStatusRequest := todo.TodoStatusUpdateRequest{}
		err = c.BodyParser(&todoStatusRequest)
		if err != nil {
			log.Fatal(err)
		}

		service.UpdateTodoStatus(uint(todoID), todoStatusRequest)
		return c.SendStatus(fiber.StatusOK)
	}
}
