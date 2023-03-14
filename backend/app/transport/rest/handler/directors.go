package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DirectorHandler struct {
	service DirectorService
}

func NewDirectorHandler(s DirectorService) DirectorHandler {
	return DirectorHandler{service: s}
}

func (h *DirectorHandler) create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"create": "res"})
}

func (h *DirectorHandler) get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"get": "res"})
}
