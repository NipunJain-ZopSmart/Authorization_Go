package controllers

import "example.com/models"

type Userstore interface {
	Register(username, email, password, role string) error
	CheckUser(email string) (*models.User, error)
}
