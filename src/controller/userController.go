package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohandas-max/ambassador/src/database"
	"github.com/rohandas-max/ambassador/src/model"
)

func Ambassador(c *fiber.Ctx) error {
	var users []model.User

	database.DB.Where("is_ambassador = ?", true).Find(&users)
	return c.JSON(users)
}
