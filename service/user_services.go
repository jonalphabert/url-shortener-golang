package service

import (
	"errors"

	"github.com/jonalphabert/url-shortener-golang/logger"
	"github.com/jonalphabert/url-shortener-golang/models"
	"github.com/jonalphabert/url-shortener-golang/repository"
)

var ErrInvalidInput = errors.New("invalid input")

type UserService struct {
    repo repository.UserRepository
    log  *logger.LoggerType // sesuaikan dengan tipe loggermu; bisa *logrus.Logger
}

func NewUserService(repo repository.UserRepository, log *logger.LoggerType) *UserService {
    return &UserService{repo: repo, log: log}
}

func (s *UserService) CreateUser(name string) (*models.User, error) {
    if name == "" {
        return nil, ErrInvalidInput
    }
    u := &models.User{Name: name}
    return s.repo.Create(u)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
    return s.repo.GetAll()
}

func (s *UserService) GetUser(id int) (*models.User, error) {
    return s.repo.GetByID(id)
}

func (s *UserService) DeleteUser(id int) (*models.User, error) {
    return s.repo.Delete(id)
}
