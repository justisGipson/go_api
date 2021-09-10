package routes

import "github.com/gofiber/fiber/v2"

func RouteNotFound(a *fiber.App) {
	// register new route for "Page not found"
	a.Use(
		func(c *fiber.Ctx) error {
			// return 404 status + json res
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "Page not found",
			})
		},
	)
}
