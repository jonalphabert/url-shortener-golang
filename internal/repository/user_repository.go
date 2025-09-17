package repository

import (
	"github.com/jonalphabert/url-shortener-golang/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (res *models.User, err error) {
    if err := r.db.Create(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
    var users []*models.User
    if err := r.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
    var user models.User
    if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Delete(id int) (*models.User, error) {
    var user models.User
    if err := r.db.Where("id = ?", id).Delete(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}