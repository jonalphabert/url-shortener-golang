package repository

import "github.com/jonalphabert/url-shortener-golang/models"

type UserRepository interface {
    Create(u *models.User) (*models.User, error)
    GetByID(id int) (*models.User, error)
    GetAll() ([]*models.User, error)
    Delete(id int) (*models.User, error)
}
