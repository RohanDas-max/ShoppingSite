package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/rohandas-max/ambassador/src/database"
	"github.com/rohandas-max/ambassador/src/model"
)

func PopulateProduct() {
	database.Connection()
	for i := 0; i <= 30; i++ {
		product := model.Product{
			Id:          0,
			Title:       faker.Username(),
			Description: faker.Paragraph(),
			Image:       faker.URL(),
			Price:       float64(faker.RandomUnixTime()),
		}
		database.DB.Create(&product)
	}
}
