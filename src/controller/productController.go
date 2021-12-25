package controller

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

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

func ProductsFE(c *fiber.Ctx) error {
	var products []model.Product

	ctx := context.Background()

	res, err := database.Cache.Get(ctx, "products_frontend").Result()

	if err != nil {
		database.DB.Find(&products)
		bytes, err := json.Marshal(products)

		if err != nil {
			panic(err)
		}
		if errKey := database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}

	} else {
		json.Unmarshal([]byte(res), &products)
	}
	return c.JSON(products)
}

func ProductsBE(c *fiber.Ctx) error {
	var products []model.Product

	ctx := context.Background()

	res, err := database.Cache.Get(ctx, "products_backend").Result()

	if err != nil {
		database.DB.Find(&products)
		bytes, err := json.Marshal(products)

		if err != nil {
			panic(err)
		}
		database.Cache.Set(ctx, "products_backend", bytes, 30*time.Minute)

	} else {
		json.Unmarshal([]byte(res), &products)
	}
	var searchedProducts []model.Product
	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			if strings.Contains(product.Title, lower) || strings.Contains(product.Description, lower) {
				searchedProducts = append(searchedProducts, product)
			}
		}
	} else {
		searchedProducts = products
	}
	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price < searchedProducts[j].Price
			})
		} else if sortLower == "dsc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price > searchedProducts[j].Price
			})
		}
	}
	total := len(searchedProducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	ItemsPerPage := 9
	var data []model.Product

	if total <= page*ItemsPerPage && total >= (page-1)*ItemsPerPage {
		data = searchedProducts[(page-1)*ItemsPerPage : total]
	} else if total >= page*ItemsPerPage {
		data = searchedProducts[(page-1)*ItemsPerPage : page*ItemsPerPage]
	} else {
		return c.JSON(fiber.Map{
			"error": "page not found",
		})
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/ItemsPerPage + 1,
	})
}
