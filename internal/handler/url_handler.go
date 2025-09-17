package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jonalphabert/url-shortener-golang/internal/logger"
	"github.com/jonalphabert/url-shortener-golang/internal/service"
)

type UrlHandler struct {
	svc *service.UrlService
	log *logger.LoggerType
}

func NewUrlHandler(svc *service.UrlService, log *logger.LoggerType) *UrlHandler {
	return &UrlHandler{svc: svc, log: log}
}

func (h *UrlHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/urls", h.GetAll)
	rg.GET("/urls/:id", h.GetByID)
	rg.POST("/urls", h.Create)
	rg.DELETE("/urls/:id", h.Delete)
	rg.PATCH("/urls/:id", h.Update)
}

func (h *UrlHandler) RegisterRedirectRoutes(rg *gin.RouterGroup) {
	rg.GET("/s/:shortUrl", h.Redirect)
}

func (h *UrlHandler) GetAll(c *gin.Context) {
	urls, err := h.svc.GetAllUrls()
	if err != nil {
		h.log.WithError(err).Error("GetAllUrls failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"urls": urls})
}

func (h *UrlHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	url, err := h.svc.GetUrl(id)
	if err != nil {
		h.log.WithField("id", id).Warn("Url not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *UrlHandler) Create(c *gin.Context) {
	var body struct {
		UserID 		int `json:"user_id" binding:"required"` // required field	uint
		ShortUrl 	string `json:"short_url" default:""` // required field
		LongUrl  	string `json:"long_url" binding:"required"`   // required field
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	url, err := h.svc.CreateUrl(body.UserID, body.ShortUrl, body.LongUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"url": url})
}

func (h *UrlHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.svc.DeleteUrl(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Url deleted"})
}

func (h *UrlHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var body struct {
		ShortUrl string `json:"short_url" binding:"required"` // required field
		LongUrl  string `json:"long_url" binding:"required"`   // required field
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	url, err := h.svc.UpdateUrl(id, body.ShortUrl, body.LongUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *UrlHandler) Redirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	url, err := h.svc.GetUrlByShortUrl(shortUrl)
	if err != nil {
		h.log.WithField("shortUrl", shortUrl).Warn("Url not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.LongUrl)
}