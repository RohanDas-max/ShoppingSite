package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"-"`
}

func (user *User) SetPass(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePass(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
