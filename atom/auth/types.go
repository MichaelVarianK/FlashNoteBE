package atom_auth

import (
	"time"
)

type User struct {
	UserId      uint           `gorm:"primaryKey" json:"user_id"`
	FullName    string         `gorm:"not null" json:"full_name"`
	Email       string         `gorm:"not null" json:"email"`
	Password    string         `gorm:"not null" json:"password"`
	PhoneNumber string         `json:"phone_number"`
	CreatedDate time.Time      `json:"created_date"`
	Provider    string         `json:"provider"`
	Tokens      []RefreshToken `gorm:"foreignKey:UserId;references:UserId" json:"tokens,omitempty"`
}

type RefreshToken struct {
	RefreshTokenId uint      `gorm:"primaryKey" json:"refresh_token_id"`
	UserId         uint      `json:"user_id"`
	RefreshToken   string    `gorm:"not null" json:"refresh_token"`
	CreatedDate    time.Time `json:"created_date"`
	ExpirationDate time.Time `json:"expiration_date"`
}

// Request bodies

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Response bodies

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GoogleAuthStateData struct {
	State string `json:"state"`
	URL   string `json:"url"`
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}
