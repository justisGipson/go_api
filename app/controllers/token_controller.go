package controllers

import (
	"github.com/CodeliciousProduct/bluebird/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetNewAccessToken
// @Description Create new access token.
// @Summary create new access token
// @Tags token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	// generate new token
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})

}
