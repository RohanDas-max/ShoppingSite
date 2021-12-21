package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("Secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"msg": "unauthanticated",
		})
	}
	return c.Next()
}

func GetUserID(c *fiber.Ctx) (int, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("Secret"), nil
	})
	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.Atoi(payload.Subject)
	return id, nil
}
