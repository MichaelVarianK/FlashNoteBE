package atom_auth

import "time"

type User struct {
	UserId uint `gorm:"primaryKey" json:"user_id"`
	FullName string `gorm:"not null" json:"full_name"`
	Email string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedDate time.Time `json:"created_date"`
}