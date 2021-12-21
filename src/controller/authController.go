package controller

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rohandas-max/ambassador/src/database"
	"github.com/rohandas-max/ambassador/src/middlewares"
	"github.com/rohandas-max/ambassador/src/model"
)

//**REGISTER**
func Register(c *fiber.Ctx) error {

	var data = make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"opps": "password do not match",
		})
	}

	user := model.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: false,
	}
	user.SetPass(data["password"])
	database.DB.Create(&user)

	return c.JSON(user)
}

//**LOGIN**
func Login(c *fiber.Ctx) error {
	var data = make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user model.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "invalid Credentials",
		})
	}
	if err := user.ComparePass(data["password"]); err != nil {
		return c.JSON(fiber.Map{
			"eeehhh!!": "invalid Credentials",
		})
	}

	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("Secret"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Invalid Credentials",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"msg": "success",
	})
}

//**LOGOUT**
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"msg": "log out success",
	})
}

//**USER**
func User(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserID(c)

	var user model.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

//**UPDATE INFO**
func UpdateInfo(c *fiber.Ctx) error {
	var data = make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	id, _ := middlewares.GetUserID(c)

	user := model.User{
		Id:        id,
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	database.DB.Model(model.User{}).Where("Id =?", id).Updates(&user)
	return c.JSON(&user)
}

//**UPDATE PASSWORD**
func UpdatePassword(c *fiber.Ctx) error {

	var data = make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"opps": "password do not match",
		})
	}
	id, _ := middlewares.GetUserID(c)

	user := model.User{
		Id: id,
	}
	user.SetPass(data["password"])
	database.DB.Model(model.User{}).Where("Id =?", id).Updates(&user)
	return c.JSON(fiber.Map{
		"msg": "password changed",
	})

}
