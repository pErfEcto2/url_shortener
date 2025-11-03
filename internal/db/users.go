package db

import (
	"errors"
	"github.com/pErfEcto2/url_shortener/internal/models"
)

var users []models.User

func AddUser(user models.User) error {
	for _, u := range users {
		if u.Username == user.Username {
			return errors.New("user already exists")
		}
	}

	if user.Username == "" || user.Password == "" {
		return errors.New("invalid user")
	}

	users = append(users, user)

	return nil
}

func GetUsers() []models.User {
	return users
}

func GetUser(user models.User) models.User {
	for _, u := range users {
		if u.Username == user.Username {
			return u
		}
	}
	return models.User{}

}

func HasUser(user models.User) bool {
	for _, u := range users {
		if u.Username == user.Username {
			return true
		}
	}
	return false
}
