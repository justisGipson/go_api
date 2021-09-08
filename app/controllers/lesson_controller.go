package controllers

import (
	"fmt"

	"github.com/CodeliciousProduct/bluebird/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetLessons func gets all existing lessons
// @Description get all existing lessons
// @Summary get all lessons
// @Tags lessons
// @Accept json
// @Produce json
// @Success 200 {array} models.Lesson
// @Router /v1/lessons [get]
func GetLessons(c *fiber.Ctx) error {
	// create dn conn
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// get all lessons
	lessons, err := db.GetLessons()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     fmt.Errorf("lessons not found"),
			"count":   0,
			"lessona": nil,
		})
	}
	// return 200 "OK"
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(lessons),
		"lessons": lessons,
	})
}

// GetLesson func gets lesson by given ID or returns 404 error.
// @Description Get lesson by given ID.
// @Summary get lesson by given ID
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Success 200 {object} models.Lesson
// @Router /v1/lesson/{id} [get]
func GetLesson(c *fiber.Ctx) error {
	// grab lesson by id in url
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// create db conn
	db, err := database.OpenDBConnection()
	if err != nil {
		// return 500 status & db conn error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// get lesson by id
	lesson, err := db.GetLesson(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    fmt.Errorf("lesson %l not found", lesson),
			"lesson": nil,
		})
	}
	// return 200 "OK"
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"lesson": lesson,
	})
}

// CreateLesson func
