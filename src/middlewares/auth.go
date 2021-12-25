package middlewares

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const SK = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SK), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"msg": "unauthanticated",
		})
	}

	payload := token.Claims.(*ClaimsWithScope)
	IsAmbassador := strings.Contains(c.Path(), "/api/ambassador")
	if payload.Scope == "admin" && IsAmbassador || payload.Scope == "ambassador" && !IsAmbassador {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"msg": "unauthorized",
		})
	}

	return c.Next()

}

func GenerateJWT(id uint, scope string) (string, error) {
	payload := ClaimsWithScope{}
	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	payload.Scope = scope
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SK))
}

func GetUserID(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SK), nil
	})
	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)
	id, _ := strconv.Atoi(payload.Subject)
	return uint(id), nil
}
