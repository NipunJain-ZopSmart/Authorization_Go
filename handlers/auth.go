package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type authHandler struct {
	authcontroller AuthController
}

func NewAuthhandler(authcontroller AuthController) *authHandler {
	return &authHandler{authcontroller: authcontroller}
}
func (handler *authHandler) Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	name := data["name"]
	password := data["password"]
	email := data["email"]
	role := data["role"]

	err := handler.authcontroller.Register(name, email, password, role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"error":   nil,
		"message": "Registration Success",
	})
}

func (handler *authHandler) Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
			"message": "Error in Parsing the request body",
		})
	}
	password := data["password"]
	email := data["email"]
	userData, err := handler.authcontroller.Login(email, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
			"message": "Error in Login",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"data": userData,
	})
}

func (handler *authHandler) GetAdminDetails(c *fiber.Ctx) error {
	userCredentials := c.Locals("userCredentials").(jwt.MapClaims)

	user, err := handler.authcontroller.GetAdminDetails(userCredentials["email"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})

}
