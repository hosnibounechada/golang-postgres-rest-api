package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/hosnibounechada/go-api/internal/auth/models"
	"github.com/hosnibounechada/go-api/internal/auth/util"
	database "github.com/hosnibounechada/go-api/internal/db"
)

type AuthService struct {
	// Implement user-related service methods here.
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(payload models.RegisterUserReq) (models.RegisterUserRes, error) {

	// Check if the username (email) is already taken
	existingUser := models.User{}
	if !database.DB.Where("email = ?", payload.Email).First(&existingUser).RecordNotFound() {
		return models.RegisterUserRes{}, errors.New("Email already exists")
	}

	hashedPassword, err := util.HashPassword(payload.Password)
	if err != nil {
		// Handle the error appropriately
		return models.RegisterUserRes{}, err
	}

	username := util.GenerateRandomUsername(payload.FirstName, payload.LastName)

	newUser := models.User{
		FirstName: strings.ToLower(payload.FirstName),
		LastName:  strings.ToLower(payload.LastName),
		Username:  username,
		Email:     payload.Email,
		Password:  string(hashedPassword),
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		// Handle the error appropriately
		return models.RegisterUserRes{}, err
	}

	createUserRes := models.RegisterUserRes{
		ID:        newUser.ID,
		Username:  newUser.Username,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}

	// Link should be sent in email
	refreshToken, _ := util.GenerateConfirmationToken(newUser.ID)
	fmt.Println(refreshToken)

	return createUserRes, nil
}

func (s *AuthService) AccountConfirmation(userId int64) {

	user := models.User{}

	database.DB.Where("id = ?", userId).First(&user)

	user.Verified = true

	database.DB.Save(&user)
}

func (s *AuthService) Login(payload models.LoginUserReq) (models.LoginUserRes, error) {
	existingUser := models.User{}

	if database.DB.Where("email = ?", payload.Email).First(&existingUser).RecordNotFound() {
		return models.LoginUserRes{}, errors.New("Invalid Credentials")
	}

	if err := util.CheckPassword(payload.Password, existingUser.Password); err != nil {
		return models.LoginUserRes{}, errors.New("Invalid Credentials")
	}

	if !existingUser.Verified {
		return models.LoginUserRes{}, errors.New("Account isn't verified yet!")
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return models.LoginUserRes{}, errors.New("Error generating UUID")
	}
	uuidStr := u.String()

	accessToken, _ := util.GenerateAccessToken(existingUser)
	refreshToken, _ := util.GenerateRefreshToken(existingUser.ID, uuidStr)

	if err != nil {
		return models.LoginUserRes{}, errors.New("Failed to complete token process")
	}

	jwt := models.TokenRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	newDevice := models.Device{
		Name:    "Lenovo ThinkPad",
		OS:      "Widows",
		Browser: "Chrome",
		Token:   uuidStr,
		UserID:  existingUser.ID,
	}

	if err := database.DB.Create(&newDevice).Error; err != nil {
		return models.LoginUserRes{}, err
	}

	loginRes := models.LoginUserRes{
		User: existingUser,
		Jwt:  jwt,
	}

	return loginRes, nil
}

func (s *AuthService) Logout(userId int64, token string) {

	database.DB.Where("user_id = ? AND token = ?", userId, token).Delete(&models.Device{})
}

func (s *AuthService) Refresh(userId int64, token string) (string, error) {
	user := models.User{}
	device := models.Device{}

	if database.DB.Where("id = ?", userId).First(&user).RecordNotFound() || database.DB.Where("user_id = ? AND token = ?", userId, token).First(&device).RecordNotFound() {
		return "", errors.New("Invalid User or Token")
	}

	if token != device.Token {
		return "", errors.New("Invalid Token")
	}

	accessToken, _ := util.GenerateAccessToken(user)

	return accessToken, nil
}
