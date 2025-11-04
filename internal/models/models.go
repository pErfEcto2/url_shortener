package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string            `form:"username" json:"username"`
	Password string            `form:"password" json:"password"`
	Urls     map[string]string // map of shortened urls and originals
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(bytes)
	return err
}

func (u *User) CompareHashedPasswords(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
