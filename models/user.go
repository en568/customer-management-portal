package models

import (
	"enormous-air-portal/config"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	err := config.DB.QueryRow("SELECT id, name, email, password, role FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
