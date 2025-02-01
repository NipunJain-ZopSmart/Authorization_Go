package store

import (
	"database/sql"
	"example.com/models"
)

type userStore struct {
	db *sql.DB
}

func (u *userStore) GetAdminDetails(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserStore(db *sql.DB) *userStore {
	return &userStore{db: db}
}

func (u *userStore) Register(username, email, password, role string) error {
	query := `INSERT INTO User (name, email, password,role) VALUES (?, ?, ?, ?)`
	_, err := u.db.Exec(query, username, email, password, role)
	if err != nil {
		return err
	}
	return nil
}

func (u *userStore) CheckUser(email string) (*models.User, error) {
	query := "SELECT * FROM User WHERE email = ?"
	row := u.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
