package service

import (
	"errors"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/jonalphabert/url-shortener-golang/internal/logger"
	"github.com/jonalphabert/url-shortener-golang/internal/models"
	"github.com/jonalphabert/url-shortener-golang/internal/repository"
)

type UrlService struct {
	repo repository.UrlRepository
	log  *logger.LoggerType
}

var ErrShortUrlExists = errors.New("short url already exists")

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func NewUrlService(repo *repository.UrlRepository, log *logger.LoggerType) *UrlService {
	return &UrlService{repo: *repo, log: log}
}

func (s *UrlService) CreateUrl(ownerId int,shortUrl string, longUrl string) (*models.Url, error) {
	if longUrl == "" {
		s.log.Error("long url is empty")
		return nil, ErrInvalidInput
	}
	
	parsedURL, err := url.Parse(longUrl)
	if err != nil {
		s.log.Error("long url is invalid")
		return nil, ErrInvalidInput
	}

	// Validasi URL lebih ketat
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		s.log.Error("URL must have http or https scheme")
		return nil, ErrInvalidInput
	}

	if parsedURL.Host == "" {
		s.log.Error("URL must have a valid host")
		return nil, ErrInvalidInput
	}

	// Pastikan host memiliki domain yang valid (mengandung titik)
	if !strings.Contains(parsedURL.Host, ".") {
		s.log.Error("URL must have a valid domain")
		return nil, ErrInvalidInput
	}

	if shortUrl == "" {
		// Generate random short URL jika tidak disediakan
		for {
			shortUrl = generateRandomString(6)
			// Cek apakah sudah ada
			if _, err := s.repo.GetByShortUrl(shortUrl); err != nil {
				break // URL belum ada, bisa dipakai
			}
		}
	} else {
		// Cek apakah shortUrl yang diberikan sudah ada
		if _, err := s.repo.GetByShortUrl(shortUrl); err == nil {
			s.log.Error("short url already exists")
			return nil, ErrShortUrlExists
		}
	}


	u := &models.Url{UserID: int(ownerId), ShortUrl: shortUrl, LongUrl: longUrl}
	return s.repo.Create(u)
}

func (s *UrlService) GetAllUrls() ([]*models.Url, error) {
	return s.repo.GetAll()
}

func (s *UrlService) GetUrl(id int) (*models.Url, error) {
	return s.repo.GetByID(id)
}

func (s *UrlService) DeleteUrl(id int) (error) {
	return s.repo.DeleteUrl(id)
}

func (s *UrlService) UpdateUrl(id int, shortUrl string, longUrl string) (*models.Url, error) {
	if shortUrl == "" || longUrl == "" {
		return nil, ErrInvalidInput
	}
	u := &models.Url{ShortUrl: shortUrl, LongUrl: longUrl}
	return s.repo.UpdateUrl(id, u)
}

func (s *UrlService) GetUrlByShortUrl(shortUrl string) (*models.Url, error) {
	return s.repo.GetByShortUrl(shortUrl)
}

func (s *UrlService) GetUserUrls(userid int) ([]*models.Url, error) {
	return s.repo.GetUserUrls(userid)
}