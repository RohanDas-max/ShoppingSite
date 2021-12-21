package model

type Link struct {
	Id      uint      `json:"id"`
	Code    string    `json:"code"`
	UserId  uint      `json:"user_id"`
	User    User      `json:"user" gorm:"foreignKey:UserId"`
	Product []Product `json:"products" gorm:"many2many:link_products"`
	Order   Order     `json:"order" gorm:"-"`
}
