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

// GetCourses func gets all existing courses
// @Description get all existing courses
// @Summary get all courses
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {array} models.Course
// @Router /v1/courses [get]
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

// GetCourse func gets course by given ID or returns 404 error.
// @Description Get course by given ID.
// @Summary get course by given ID
// @Tags Course
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} models.Course
// @Router /v1/course/{id} [get]
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

// CreateCourse func to add new course
// @Description Create new course
// @Summary create new course
// @Tags Course
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param courseNumber body string true "Course Number"
// @Success 200 {object} models.Course
// @Security ApiKeyAuth
// @Router /v1/course [post]
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

// UpdateCourse to update by id
// @Description update course with a given id
// @Summary update course by id
// @Tags Course
// @Accept json
// @Produce json
// @Param id body string true "Course ID"
// @Param name body string true "Course Name"
// @Param courseNumber body string true "Course Number"
// @Param active body bool true "Course Active status"
// @Param course_attrs body models.CourseAttrs true "Course Attributes"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/course [put]
func UpdateCourse(c *fiber.Ctx) error {
	now := time.Now().Unix()
	// jwt claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// set expiration of jwt token from course jwt
	expires := claims.Expires
	// check jwt claim expiration
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized - check token, could be expired",
		})
	}
	course := &models.Course{}
	if err := c.BodyParser(course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
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
	// validate course with given ID
	courseToUpdate, err := db.GetCourse(course.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "course not found",
		})
	}
	// set default course data
	course.Updated_at = time.Now() // Unix() here too?
	// course validator
	validate := utils.NewValidator()
	// course fields validation
	if err := validate.Struct(course); err != nil {
		// return error if some or all fields are invalid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	// update course entry
	if _, err := db.UpdateCourse(courseToUpdate.ID, course); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusCreated)
}

// DeleteCourse to delete course by ID
// @Description delete course with given ID
// @Summary delete course by id
// @Tags Course
// @Accept json
// @Produce json
// @Param id body string "Course ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/course [delete]
func DeleteCourse(c *fiber.Ctx) error {
	now := time.Now().Unix()
	// get jwt claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// claims expiration set on current course
	expires := claims.Expires
	// check claim expiration against time.Now
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized - token is expired",
		})
	}
	// new course struct
	course := &models.Course{}
	// validate json
	if err := c.BodyParser(course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// new course model validator
	validate := utils.NewValidator()
	// validate course id field
	if err := validate.StructPartial(course, "id"); err != nil {
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
	// verify course exists
	courseToDel, err := db.GetCourse(course.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   fmt.Errorf("course %c not found - %v", course.ID, err),
		})
	}
	// delete course with given ID
	if _, err := db.DeleteCourse(courseToDel.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
