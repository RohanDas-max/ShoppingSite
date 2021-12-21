package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/rohandas-max/ambassador/src/database"
	"github.com/rohandas-max/ambassador/src/model"
)

func PopulateUser() {
	database.Connection()
	for i := 0; i <= 30; i++ {
		ambassador := model.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		ambassador.SetPass("1234")
		database.DB.Create(&ambassador)
	}
}
