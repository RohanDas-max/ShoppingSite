package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id           uint     `json:"id"`
	FirstName    string   `json:"firstname"`
	LastName     string   `json:"lastname"`
	Email        string   `json:"email" gorm:"unique"`
	Password     []byte   `json:"-"`
	IsAmbassador bool     `json:"-"`
	Revenue      *float64 `json:"revenue,omitempty" gorm:"-"`
}

func (user *User) SetPass(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePass(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

type Admin User
type Ambassador User

func (admin *Admin) CalculateRevenue(db *gorm.DB) {
	var order []Order
	db.Preload("OrderItems").Find(&order, &Order{
		UserId:   admin.Id,
		Complete: true,
	})

	var revenue float64 = 0
	for _, order := range order {
		for _, OrderItem := range order.OrderItems {
			revenue += OrderItem.AdminRevenue
		}
	}
	admin.Revenue = &revenue
}

func (ambassador *Ambassador) CalculateRevenue(db *gorm.DB) {
	var order []Order
	db.Preload("OrderItems").Find(&order, &Order{
		UserId:   ambassador.Id,
		Complete: true,
	})

	var revenue float64 = 0
	for _, order := range order {
		for _, OrderItem := range order.OrderItems {
			revenue += OrderItem.AmbassadorRevenue
		}
	}
	ambassador.Revenue = &revenue
}
