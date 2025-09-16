package repository

import "github.com/jonalphabert/url-shortener-golang/models"

type UrlRepository interface {
    Create(u *models.Url) (*models.Url, error)
    GetByID(id int) (*models.Url, error)
	GetByShortUrl(shortUrl string) (*models.Url, error)
    GetAll() ([]*models.Url, error)
    Delete(id int) (*models.Url, error)
	Update(id int, u *models.Url) (*models.Url, error)
}