package atom_auth

import (
	"github.com/KelXR/FlashNoteBE/config"
)

func CreateUserDB(user *User) error {
	db := config.DB
	return db.Create(user).Error
}

func GetUserByEmailDB(email string, provider string) (User, error) {
	db := config.DB
	var user User
	err := db.Where("email = ? AND provider = ?", email, provider).First(&user).Error
	return user, err
}

func SaveRefreshTokenDB(token *RefreshToken) error {
	db := config.DB
	return db.Create(token).Error
}

func GetRefreshTokenDB(tokenStr string) (RefreshToken, User, error) {
	db := config.DB
	var token RefreshToken
	var user User
	err := db.Preload("User").Where("refresh_token = ?", tokenStr).First(&token).Error
	if err != nil {
		return token, user, err
	}
	err = db.Model(&token).Association("User").Find(&user)
	return token, user, err
}

func DeleteRefreshTokenDB(tokenStr string) {
	db := config.DB
	db.Where("refresh_token = ?", tokenStr).Delete(&RefreshToken{})
}
