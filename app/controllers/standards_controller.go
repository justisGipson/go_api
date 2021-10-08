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

// GetStandards gets all standards
// @Description get all current standards
// @Summary get all standards
// @Tags standards
// @Accept json
// @Produce json
// @Param Grade/list of grades?
// @Success 200 {array} models.Standard
// @Router /v1/standards [get]
func GetStandards(c *fiber.Ctx) error {
	// create db connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// get all standards
	standards, err := db.GetStandards()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":     true,
			"msg":       fmt.Errorf("standards not found - %v", err),
			"count":     0,
			"standards": nil,
		})
	}
	// return 200 "OK"
	return c.JSON(fiber.Map{
		"error":     false,
		"msg":       nil,
		"count":     len(standards),
		"standards": standards,
	})
}

// GetStandard get single standard
// @Description return single standard
// @Summary single standard
// @Tags Standards
// @Accept json
// @Produce json
// @Param grade/list of grades?
// @Param id
// @Success 200 {object} models.Standard
// @Router /v1/standards/{id} [get]
func GetStandard(c *fiber.Ctx) error {
	// fetch standard w/ id
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
		// return 500 and db connection error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// get standard w/ id
	standard, err := db.GetStandard(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      fmt.Errorf("standard %c not found - %v", id, err),
			"standard": nil,
		})
	}
	// return 200 "OK"
	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil, // success message?
		"standard": standard,
	})
}

// UpdateStandard updates standard by id
// @Description updates standard by id
// @Summary updates standard by id
// @Tags Standard
// @Accept json
// @Produce json
// @Param id
// @Param state
// @Param org
// @Param ... could be more, have to think about what "could" be updated
// @Success 201 {string} status "OK"
// @Router /v1/standards [put]
func UpdateStandard(c *fiber.Ctx) error {
	// get time of update
	now := time.Now().Unix()
	// jwt claims
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// set jwt expiration from standard jwt data
	expires := claims.Expires
	// if time.Now > jwt claim expiration
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "401 - unauthorized; token is expired",
		})
	}
	standard := &models.Standard{}
	if err := c.BodyParser(standard); err != nil {
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
	// validate standard with id
	standardToUpdate, err := db.GetStandard(standard.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   fmt.Errorf("standard %c not found %v", standard.ID, err),
		})
	}
	// set field to be updated for standard
	standard.Updated_at = time.Now() // Unix()?
	// standard validator
	validate := utils.NewValidator()
	// standard fields validation
	if err := validate.Struct(standard); err != nil {
		// return error if some/all fields are invalid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	// update standard entry
	if _, err := db.UpdateStandard(standardToUpdate.ID, standard); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusCreated)
}
