package repository

import (
	"github.com/jonalphabert/url-shortener-golang/internal/models"
	"gorm.io/gorm"
)

type UrlRepository struct {
    db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
    return &UrlRepository{db: db}
}

func (r *UrlRepository) Create(url *models.Url) (res *models.Url, err error) {
    if err := r.db.Create(url).Error; err != nil {
        return nil, err
    }
    return url, nil
}

func (r *UrlRepository) GetAll() ([]*models.Url, error) {
    var urls []*models.Url
    if err := r.db.Find(&urls).Error; err != nil {
        return nil, err
    }
    return urls, nil
}

func (r *UrlRepository) GetByShortUrl(shortUrl string) (*models.Url, error) {
    var url models.Url
    if err := r.db.Where("short_url = ?", shortUrl).First(&url).Error; err != nil {
        return nil, err
    }
    return &url, nil
}

func (r *UrlRepository) GetByID(id int) (*models.Url, error) {
    var url models.Url
    if err := r.db.Where("id = ?", id).First(&url).Error; err != nil {
        return nil, err
    }
    return &url, nil
}

func (r *UrlRepository) UpdateUrl(id int, url *models.Url) (*models.Url, error) {
    if err := r.db.Save(url).Error; err != nil {
        return nil, err
    }
    return url, nil
}

func (r *UrlRepository) DeleteUrl(id int) error {
    if err := r.db.Delete(&models.Url{}, id).Error; err != nil {
        return err
    }
    return nil
}