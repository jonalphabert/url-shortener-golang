package repository

import "github.com/jonalphabert/url-shortener-golang/internal/models"

type UserRepository interface {
    Create(u *models.UserInMemory) (*models.UserInMemory, error)
    GetByID(id int) (*models.UserInMemory, error)
    GetAll() ([]*models.UserInMemory, error)
    Delete(id int) (*models.UserInMemory, error)
}
