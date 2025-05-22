package atom_auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/KelXR/FlashNoteBE/config"
	utils_google "github.com/KelXR/FlashNoteBE/utils/google"
	utils_token "github.com/KelXR/FlashNoteBE/utils/token"
)

func CreateUserUseCase(user RegisterInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := User{
		Email:       user.Email,
		FullName:    user.Username,
		Password:    string(hashedPassword),
		CreatedDate: time.Now(),
		Provider:    "local",
	}

	return CreateUserDB(&newUser)
}

func LoginUserUseCase(user LoginInput) (Tokens, error) {
	getUser, err := GetUserByEmailDB(user.Email, "local")
	if err != nil {
		log.Print("[atom][auth][resource.go][LoginUserUseCase] Error: ", err)
		return Tokens{}, errors.New("User not found")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(user.Password))
	if err != nil {
		log.Print("[atom][auth][resource.go][LoginUserUseCase] Error: ", err)
		return Tokens{}, errors.New("Invalid password")
	}

	// Generate access and refresh token
	accessToken, err := utils_token.GenerateAccessToken(user.Email)
	if err != nil {
		log.Print("[atom][auth][resource.go][LoginUserUseCase] Error: ", err)
		return Tokens{}, err
	}

	refreshToken := utils_token.GenerateRefreshToken()

	err = SaveRefreshTokenDB(&RefreshToken{
		UserId:         getUser.UserId,
		RefreshToken:   refreshToken,
		CreatedDate:    time.Now(),
		ExpirationDate: time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Print("[atom][auth][resource.go][LoginUserUseCase] Error: ", err)
		return Tokens{}, err
	}

	tokens := Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func RefreshAuthUseCase(refreshToken RefreshTokenInput) (Tokens, error) {
	getRefreshToken, getUser, err := GetRefreshTokenDB(refreshToken.RefreshToken)
	if err != nil || time.Now().After(getRefreshToken.ExpirationDate) {
		log.Print("[atom][auth][resource.go][RefreshAuthUseCase] Error: ", err)
		return Tokens{}, errors.New("Invalid or expired refresh token")
	}

	DeleteRefreshTokenDB(getRefreshToken.RefreshToken)

	newAccessToken, err := utils_token.GenerateAccessToken(getUser.Email)
	if err != nil {
		log.Print("[atom][auth][resource.go][RefreshAuthUseCase] Error: ", err)
		return Tokens{}, err
	}

	newRefreshToken := utils_token.GenerateRefreshToken()

	err = SaveRefreshTokenDB(&RefreshToken{
		UserId:         getUser.UserId,
		RefreshToken:   newRefreshToken,
		CreatedDate:    time.Now(),
		ExpirationDate: time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Print("[atom][auth][resource.go][RefreshAuthUseCase] Error: ", err)
		return Tokens{}, err
	}

	tokens := Tokens{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	return tokens, nil
}

// Google Auth

func GoogleLoginUserCase() (string, string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Print("[atom][auth][resource.go][GoogleLoginUserCase] Error: ", err)
		return "", "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	authURL := config.GoogleAuthConfig.AuthCodeURL(state)

	return state, authURL, nil
}

func GetUserInfoFromGoogle(accessToken string) (*GoogleUserInfo, error) {
	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var userInfo GoogleUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func GoogleCallbackUseCase(ctx context.Context, code string, state string, storedState string) (Tokens, error) {
	token, err := config.GoogleAuthConfig.Exchange(ctx, code)
	if err != nil {
		log.Print("[atom][auth][resource.go][GoogleCallbackUseCase] Error: ", err)
		return Tokens{}, err
	}

	userInfo, err := GetUserInfoFromGoogle(token.AccessToken)
	if err != nil {
		log.Print("[atom][auth][resource.go][GoogleCallbackUseCase] Error: ", err)
		return Tokens{}, err
	}

	getUser, err := GetUserByEmailDB(userInfo.Email, "google")
	if err != nil {
		randomPassword := utils_google.GenerateRandomPassword()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(randomPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Print("[atom][auth][resource.go][GoogleCallbackUseCase] Error: ", err)
			return Tokens{}, errors.New("Failed to hash password")
		}

		user := User{
			UserId:      getUser.UserId,
			FullName:    userInfo.Name,
			Email:       userInfo.Email,
			Password:    string(hashedPassword),
			CreatedDate: time.Now(),
			Provider:    "google",
		}
		err = CreateUserDB(&user)
		if err != nil {
			log.Print("[atom][auth][resource.go][GoogleCallbackUseCase] Error: ", err)
			return Tokens{}, errors.New("Failed to create user")
		}
		getUser = user
	}

	accessToken, err := utils_token.GenerateAccessToken(getUser.Email)
	if err != nil {
		log.Print("[atom][auth][resource.go][LoginUserUseCase] Error: ", err)
		return Tokens{}, err
	}

	refreshToken := utils_token.GenerateRefreshToken()

	err = SaveRefreshTokenDB(&RefreshToken{
		UserId:         getUser.UserId,
		RefreshToken:   refreshToken,
		CreatedDate:    time.Now(),
		ExpirationDate: time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Print("[atom][auth][resource.go][GoogleCallbackUseCase] Error: ", err)
		return Tokens{}, err
	}

	tokens := Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}
