package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte ` json:"-"`
	Role     int    `json:"role"`
}
