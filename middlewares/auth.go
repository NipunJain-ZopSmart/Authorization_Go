package middlewares

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var jwtSecret = []byte("secret")

func VerifyUser(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader { // Means "Bearer " was missing
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization format",
		})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}
	c.Locals("userCredentials", claims)
	return c.Next()
}

func CheckAdmin(c *fiber.Ctx) error {
	user := c.Locals("userCredentials").(jwt.MapClaims)

	userRole := user["role"].(string)

	if userRole != "ADMIN" {
		return errors.New("YOU ARE NOT AN ADMIN")
	}
	return c.Next()
}

func CheckUser(c *fiber.Ctx) error {
	user := c.Locals("userCredentials").(jwt.MapClaims)
	userRole := user["role"].(string)
	if userRole != "USER" {
		return errors.New("YOU ARE NOT A USER")
	}
	return c.Next()
}
