package controllers

import (
	"fmt"
	"time"

	"github.com/CodeliciousProduct/bluebird/app/models"
	"github.com/CodeliciousProduct/bluebird/internal/utils"
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
	// create db conn
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
			"msg":     fmt.Errorf("lessons not found - %v", err),
			"count":   0,
			"lessons": nil,
		})
	}
	// return 200 "OK"
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil, // do we want a success message?
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
	// grab lesson by id
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
			"msg":    fmt.Errorf("lesson %c not found - %v", id, err),
			"lesson": nil,
		})
	}
	// return 200 "OK"
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil, // do we want a success message?
		"lesson": lesson,
	})
}

// CreateLesson func to add new lesson
// @Description Create new lesson
// @Summary create new lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param lessonNumber body string true "Lesson Number"
// @Success 200 {object} models.Lesson
// @Security ApiKeyAuth
// @Router /v1/lesson [post]
func CreateLesson(c *fiber.Ctx) error {
	// get time for now
	now := time.Now().Unix()

	// jwt claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// set expiration of jwtToken from lesson jwt data
	expires := claims.Expires
	// if time.Now > jwt expire
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check token - could be expired",
		})
	}
	// create new lesson struct
	lesson := &models.Lesson{}
	// validate json data
	if err := c.BodyParser(lesson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// create db connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// new lesson model validator
	validate := utils.NewValidator()
	// set initialized default lesson data
	lesson.ID = uuid.New()
	lesson.Created_at = time.Now()
	lesson.Active = true // might need to revisit as we discuss lesson status/or old vs new/draft vs live...
	// validate lesson fields
	if err := validate.Struct(lesson); err != nil {
		// return status if some/all fields aren't valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	// create lesson entry
	if _, err := db.CreateLesson(lesson); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"lesson": lesson,
	})
}

// UpdateLesson to update by id
// @Description update lesson with a given id
// @Summary update lesson by id
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id body string true "Lesson ID"
// @Param name body string true "Lesson Name"
// @Param lessonNumber body string true "Lesson Number"
// @Param active body bool true "Lesson Active status"
// @Param lesson_attrs body models.LessonAttrs true "Lesson Attributes"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/lesson [put]
func UpdateLesson(c *fiber.Ctx) error {
	now := time.Now().Unix()

	// jwt claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// set expiration of jwtToken from lesson jwt data
	expires := claims.Expires
	// if time.Now > jwt expire
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "401 - unauthorized; token is expired",
		})
	}

	lesson := &models.Lesson{}
	if err := c.BodyParser(lesson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// create db connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// validate lesson with given ID
	lessonToUpdate, err := db.GetLesson(lesson.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   fmt.Errorf("lesson %c not found %v", lesson.ID, err),
		})
	}
	// set default lesson data
	lesson.Updated_at = time.Now()
	// lesson validator
	validate := utils.NewValidator()
	// lesson fields validation
	if err := validate.Struct(lesson); err != nil {
		// return error if some/all fields aren't valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	// update lesson entry
	if _, err := db.UpdateLesson(lessonToUpdate.ID, lesson); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DeleteLesson to delete lesson by ID
// @Description delete lesson with given ID
// @Summary delete lesson by id
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id body string "Lesson ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/lesson [delete]
func DeleteLesson(c *fiber.Ctx) error {
	now := time.Now().Unix()
	// get jwt claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// set expiration time for jwt claim on current lesson
	expires := claims.Expires
	// check time if great than jwt expiration
	if now > expires {
		// return 401 and error message
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "401 - unauthorized; token is expired",
		})
	}
	// new lesson struct
	lesson := &models.Lesson{}
	// validate json
	if err := c.BodyParser(lesson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// new validator for lesson model
	validate := utils.NewValidator()
	// validate id field on lesson
	if err := validate.StructPartial(lesson, "id"); err != nil {
		// return fields not valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	// db connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// make sure lesson exists
	lessonToDel, err := db.GetLesson(lesson.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   fmt.Errorf("lesson %c not found - %v", lesson.ID, err),
		})
	}
	// delete the lesson with ID
	if _, err := db.DeleteLesson(lessonToDel.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
