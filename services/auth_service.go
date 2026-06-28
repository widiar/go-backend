package services

import (
	"backendmaw/dto"
	"backendmaw/models"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Register(data *dto.RegisterRequest) (*dto.ResponseDto, error) {
	var count int64
	if err := s.DB.Model(&models.Users{}).Where("email = ? OR username = ?", data.Email, data.Username).Count(&count).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	if count > 0 {
		err := errors.New("user already exists")
		return new(dto.FailedResponse("User already exists", http.StatusConflict)), err
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	var user models.Users
	user.Id = uuid.NewString()
	user.Username = data.Username
	user.Email = data.Email
	user.Password = string(passwordHash)

	if err := s.DB.Create(&user).Error; err != nil {
		return new(dto.ErrorResponse()), err
	}
	dataResponse := dto.RegisterResponse{Username: user.Username, Email: user.Email}
	response := dto.SuccessResponse(&dataResponse)
	response.Status = http.StatusCreated
	return &response, nil
}

func (s *AuthService) Login(data *dto.LoginRequest) (*dto.ResponseDto, error) {
	var user models.Users
	if err := s.DB.First(&user, "username = ?", data.Username).Error; err != nil {
		var textErr string
		if errors.Is(err, gorm.ErrRecordNotFound) {
			textErr = "User not found"
		} else {
			textErr = "Internal server error"
		}
		return new(dto.FailedResponse(textErr, http.StatusUnauthorized)), err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return new(dto.FailedResponse("Password Incorrect", http.StatusUnauthorized)), err
	}
	dataToken := dto.UserToken{
		Id:    user.Id,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtKey := []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, dataToken).SignedString(jwtKey)
	if err != nil {
		return new(dto.ErrorResponse()), err
	}
	dataResp := dto.LoginResponse{
		Token: token,
	}
	return new(dto.SuccessResponse(&dataResp)), nil
}

func (s *AuthService) Me(c *echo.Context) (*dto.ResponseDto, error) {
	userToken := c.Get("user").(*dto.UserToken)
	return new(dto.SuccessResponse(&userToken)), nil
}
