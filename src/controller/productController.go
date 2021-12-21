package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rohandas-max/ambassador/src/database"
	"github.com/rohandas-max/ambassador/src/model"
)

func Products(c *fiber.Ctx) error {
	var products []model.Product

	database.DB.Find(&products)
	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.DB.Create(&product)
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product model.Product
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	// product.Id = uint(id)
	database.DB.Where("id = ?", uint(id)).Find(&product)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	product := model.Product{
		Id: uint(id),
	}
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.DB.Model(&product).Updates(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	product := model.Product{
		Id: uint(id),
	}
	database.DB.Delete(product)
	return c.JSON(fiber.Map{
		"msg": "delete success",
	})
}
