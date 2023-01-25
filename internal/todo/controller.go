package todo

import (
	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	storage *TodoStorage
}

func NewTodoController(storage *TodoStorage) *TodoController {
	return &TodoController{
		storage: storage,
	}
}

type createTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type createTodoResponse struct {
	ID string `json:"id"`
}

// @Summary Create one todo.
// @Description creates one todo.
// @Tags todos
// @Accept */*
// @Produce json
// @Param todo body createTodoRequest true "Todo to create"
// @Success 200 {object} createTodoResponse
// @Router /todos [post]
func (t *TodoController) create(c *fiber.Ctx) error {
	// parse the request body
	var req createTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the todo
	id, err := t.storage.createTodo(req.Title, req.Description, false, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createTodoResponse{
		ID: id,
	})
}

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} []todoDB
// @Router /todos [get]
func (t *TodoController) getAll(c *fiber.Ctx) error {
	// get all todos
	todos, err := t.storage.getAllTodos(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get todos",
		})
	}

	return c.JSON(todos)
}
