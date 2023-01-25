package todo

import "github.com/gofiber/fiber/v2"

func AddTodoRoutes(app *fiber.App, controller *TodoController) {
	todos := app.Group("/todos")

	// add middlewares here

	// add routes here
	todos.Post("/", controller.create)
	todos.Get("/", controller.getAll)
}
