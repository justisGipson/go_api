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
			"msg":     fmt.Errorf("courses not found - %v", err),
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

func GetCourse(c *fiber.Ctx) error {
	// grab course by id
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// get course w/ id
	course, err := db.GetCourse(id)
	if err != nil {
		// return empty course obj & error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    fmt.Errorf("course %c not found - %v", id, err),
			"course": nil,
		})
	}
	// return 200 "OK"
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil, // do we want a success message?
		"course": course,
	})
}

func CreateNewCourse(c *fiber.Ctx) error {
	// get time
	now := time.Now().Unix()

	// jwt claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// set jwt claim expiration on current course
	expires := claims.Expires
	// check time if greater than jwt expiration
	if now > expires {
		// return 401 and error msg
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized - token is expired",
		})
	}
	// new course struct
	course := &models.Course{}
	// json validation
	if err := c.BodyParser(course); err != nil {
		// return error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// create db conn
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// new course model validator
	validate := utils.NewValidator()
	// set initialized default course data
	course.ID = uuid.New()
	course.Created_at = time.Now() // Unix()?
	// may need to set other fields on creation
	if err := validate.Struct(course); err != nil {
		// return status if some/all fields are invalid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	// create course entry
	if _, err := db.CreateNewCourse(course); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil, // success msg here?
		"course": course,
	})
}
