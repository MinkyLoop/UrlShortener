package api

import (
	"net/http"
	"urlshortener/internal/service"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	shortener *service.Shortener

	redirect *service.Redirect
	baseUrl  string
}

func NewHandlers(
	shortener *service.Shortener,
	redirect *service.Redirect,
	baseUrl string,
) *Handlers {
	return &Handlers{
		shortener: shortener,
		redirect:  redirect,
		baseUrl:   baseUrl,
	}
}

func (h *Handlers) Shorten(c *gin.Context) {
	var req CreateShortenReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorRes{
			Error: "invalid request: " + err.Error(),
		})
	}

	if req.Url == "" {
		c.JSON(http.StatusBadRequest, ErrorRes{
			Error: "url is required",
		})
	}

	url, err := h.shortener.Create(c.Request.Context(), req.Url)
	if err != nil {
		c.Error(err)

		c.JSON(http.StatusInternalServerError, ErrorRes{
			Error: "failed to shorten url: " + err.Error(),
		})

		return
	}

	resp := CreateShortenRes{
		ShortUrl: h.baseUrl + "/" + url.ShortCode,
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handlers) Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, ErrorRes{
			Error: "shortCode is required",
		})

		return
	}

	originalUrl, err := h.redirect.Redirect(c.Request.Context(), shortCode)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, ErrorRes{
			Error: "failed to redirect: " + err.Error(),
		})

		return
	}

	c.Redirect(http.StatusMovedPermanently, originalUrl)
}
