package routes

import (
	"github.com/CodeliciousProduct/bluebird/app/controllers"
	"github.com/CodeliciousProduct/bluebird/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

// NOTE: ALL private routes will have full test coverage

func PrivateRoutes(a *fiber.App) {
	// routes group
	route := a.Group("/api/v1")
	// GET
	route.Get("/lesson", middleware.JWTProtected(), controllers.GetLesson)
	route.Get("lesson/:id", middleware.JWTProtected(), controllers.GetLesson)
	// POST
	route.Post("/lesson", middleware.JWTProtected(), controllers.CreateLesson)
	// PUT
	route.Put("/lesson", middleware.JWTProtected(), controllers.UpdateLesson)
	// DELETE
	route.Delete("/lesson", middleware.JWTProtected(), controllers.DeleteLesson)
}
