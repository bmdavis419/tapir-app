package todo

import "github.com/gofiber/fiber/v2"

type TodoController struct {
	storage *TodoStorage
}

func NewTodoController(storage *TodoStorage) *TodoController {
	return &TodoController{
		storage: storage,
	}
}

func (t *TodoController) create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Todo created",
	})
}
