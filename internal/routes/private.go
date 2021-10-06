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
	// Lessons routes
	route.Get("/lessons", middleware.JWTProtected(), controllers.GetLessons)
	route.Get("lessons/:id", middleware.JWTProtected(), controllers.GetLesson)
	route.Post("/lessons", middleware.JWTProtected(), controllers.CreateLesson)
	route.Put("/lessons", middleware.JWTProtected(), controllers.UpdateLesson)
	route.Delete("/lessons", middleware.JWTProtected(), controllers.DeleteLesson)
	// courses routes
	route.Get("/courses", middleware.JWTProtected(), controllers.GetCourses)
	route.Get("courses/:id", middleware.JWTProtected(), controllers.GetCourse)
	route.Post("/courses", middleware.JWTProtected(), controllers.CreateNewCourse)
	route.Put("/courses", middleware.JWTProtected(), controllers.UpdateCourse)
}
