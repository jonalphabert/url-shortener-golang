package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/jonalphabert/url-shortener-golang/internal/logger"
	"github.com/jonalphabert/url-shortener-golang/internal/models"
	"github.com/jonalphabert/url-shortener-golang/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type AuthServices struct {
	repo repository.UserRepository
	log  *logger.LoggerType
}

func NewAuthService(repo *repository.UserRepository, log *logger.LoggerType) *AuthServices {
	return &AuthServices{repo: *repo, log: log}
}
func generateJWTToken(userID uint, username string) (string, error) {
	// Create the JWT claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"username": username,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":     time.Now().Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get secret key from environment variable
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "your-256-bit-secret" // Fallback secret key
	}

	// Generate signed token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
} 

func (s *AuthServices) Login(username string, password string) (*models.Auth, error) {
	user, err := s.repo.GetUserByName(username);
	
	if err != nil {
		s.log.Error("User not found with requested username ", username, "Here is the error" , err)
		return nil, err
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := generateJWTToken(user.ID, username)
	if err != nil {
		return nil, err
	}

	return &models.Auth{Username: username, Token: token, ID: user.ID}, nil
}

func (s *AuthServices) Register(name string, password string) (*models.User, error) {
	// Check if user already exists
	_, err := s.repo.GetUserByName(name)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	s.log.Infof("Original password: %s", password)
	s.log.Infof("Hashed password: %s", string(hashedPassword))

	user, err := s.repo.Create(&models.User{Username: name, Password: string(hashedPassword)})

	if err != nil {
		return nil, err
	}

	return user, nil
}