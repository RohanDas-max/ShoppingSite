package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rohandas-max/ambassador/src/database"
	"github.com/rohandas-max/ambassador/src/model"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var links []model.Link

	database.DB.Where("user_id =? ", id).Find(&links)

	for i, link := range links {
		var orders []model.Order
		database.DB.Where("code =? and complete = true ", link.Code).Find(&orders)
		links[i].Order = orders[i]
	}
	return c.JSON(links)
}
