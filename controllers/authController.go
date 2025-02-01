package controllers

import (
	"errors"
	"example.com/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userController struct {
	userStore Userstore
}

func NewUserController(userStore Userstore) *userController {
	return &userController{
		userStore: userStore,
	}
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func (svc *userController) Register(name, email, password, role string) error {

	if name == "" || email == "" || password == "" {
		return errors.New("ALL FIELDS ARE MANDATORY")
	}
	encryptPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	err = svc.userStore.Register(name, email, encryptPassword, role)
	if err != nil {
		return err
	}
	return nil
}

func createToken(name, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"name":  name,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("Error While Creating Token")
		return "", err
	}
	return tokenString, nil
}
func (svc *userController) Login(email, password string) (map[string]interface{}, error) {
	// check if user exists
	user, err := svc.userStore.CheckUser(email)
	if err != nil {
		fmt.Println("User Doesn't Exist")
		return nil, err
	}
	// check for user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Password Doesn't Match")
		return nil, err
	}
	// create a jwt token send it to user and send user credentials except for password
	jwtToken, err := createToken(user.Name, user.Email, user.Role)
	if err != nil {
		return nil, err
	}
	userCredentials := &models.User{}
	userCredentials.Email = user.Email
	userCredentials.Name = user.Name
	return map[string]interface{}{
		"success": true,
		"user":    userCredentials,
		"token":   jwtToken,
		"message": "Login Successful",
	}, nil
}

func (svc *userController) GetAdminDetails(email string) (*models.User, error) {
	user, err := svc.userStore.CheckUser(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
