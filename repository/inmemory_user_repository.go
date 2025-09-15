package repository

import (
	"errors"
	"sync"

	"github.com/jonalphabert/url-shortener-golang/models"
)

var ErrNotFound = errors.New("user not found")

type InMemoryUserRepo struct {
    mu     sync.RWMutex
    data   map[int]*models.User
    nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{
        data:   make(map[int]*models.User),
        nextID: 1,
    }
}

func (r *InMemoryUserRepo) Create(u *models.User) (*models.User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    u.ID = r.nextID
    r.nextID++
    // copy to avoid external mutation
    copyU := *u
    r.data[u.ID] = &copyU
    return &copyU, nil
}

func (r *InMemoryUserRepo) GetByID(id int) (*models.User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    if u, ok := r.data[id]; ok {
        copyU := *u
        return &copyU, nil
    }
    return nil, ErrNotFound
}

func (r *InMemoryUserRepo) GetAll() ([]*models.User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    res := make([]*models.User, 0, len(r.data))
    for _, u := range r.data {
        copyU := *u
        res = append(res, &copyU)
    }
    return res, nil
}

func (r *InMemoryUserRepo) Delete(id int) (*models.User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    if u, ok := r.data[id]; ok {
        copyU := *u
        delete(r.data, id)
        return &copyU, nil
    }
    return nil, ErrNotFound
}
