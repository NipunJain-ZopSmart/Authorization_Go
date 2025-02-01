package handlers

import "example.com/models"

type AuthController interface {
	Register(name, email, password, role string) error
	Login(email, password string) (map[string]interface{}, error)
	GetAdminDetails(email string) (*models.User, error)
}
