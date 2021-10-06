package controllers

import (
	"fmt"

	"github.com/CodeliciousProduct/bluebird/platform/database"
	"github.com/gofiber/fiber/v2"
)

func GetCourses(c *fiber.Ctx) error {
	// create db connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// get all courses
	courses, err := db.GetCourses()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     fmt.Errorf("courses not found"),
			"count":   0,
			"courses": nil,
		})
	}
	// return "200" ok + courses obj
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil, // do we want a success message?
		"count":   len(courses),
		"courses": courses,
	})
}
