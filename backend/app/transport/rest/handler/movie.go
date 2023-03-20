package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service MovieService
}

func NewMovieHandler(s MovieService) MovieHandler {
	return MovieHandler{service: s}
}

func (h *MovieHandler) create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"create": "res"})
}

func (h *MovieHandler) get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"get": "res"})
}

func (h *MovieHandler) getAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getAll": "res"})
}
