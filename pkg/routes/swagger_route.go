package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// swagger route points to API docs route
func SwaggerRoute(a *fiber.App) {
	// routes group
	route := a.Group("/swagger")
	// route for swagger GET
	route.Get("*", swagger.Handler) // grab one user by ID
}
