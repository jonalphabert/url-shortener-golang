package repository

import (
	"errors"
	"sync"

	"github.com/jonalphabert/url-shortener-golang/models"
)

type InMemoryUrlRepo struct {
	mu sync.RWMutex
	data map[int]*models.Url
	nextID int
}

func NewInMemoryUrlRepo() *InMemoryUrlRepo {
	return &InMemoryUrlRepo{
		data: make(map[int]*models.Url),
		nextID: 1,
	}
}

func (r *InMemoryUrlRepo) Create(u *models.Url) (*models.Url, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u.ID = r.nextID
	r.nextID++
	// copy to avoid external mutation
	copyU := *u
	r.data[u.ID] = &copyU
	return &copyU, nil
}

func (r *InMemoryUrlRepo) GetByID(id int) (*models.Url, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if u, ok := r.data[id]; ok {
		copyU := *u
		return &copyU, nil
	}
	return nil, errors.New("url not found")
}

func (r *InMemoryUrlRepo) GetAll() ([]*models.Url, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res := make([]*models.Url, 0, len(r.data))
	for _, u := range r.data {
		copyU := *u
		res = append(res, &copyU)
	}
	return res, nil
}

func (r *InMemoryUrlRepo) Delete(id int) (*models.Url, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if u, ok := r.data[id]; ok {
		copyU := *u
		delete(r.data, id)
		return &copyU, nil
	}
	return nil, errors.New("url not found")
}

func (r *InMemoryUrlRepo) Update(id int, u *models.Url) (*models.Url, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; ok {
		// copy to avoid external mutation
		copyU := *u
		r.data[id] = &copyU
		return &copyU, nil
	}
	return nil, errors.New("url not found")
}

func (r *InMemoryUrlRepo) GetByShortUrl(shortUrl string) (*models.Url, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, u := range r.data {
		if u.ShortUrl == shortUrl {
			copyU := *u
			return &copyU, nil
		}
	}
	return nil, errors.New("url not found")
}