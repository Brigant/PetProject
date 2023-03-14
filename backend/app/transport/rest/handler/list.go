package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	service ListsService
}

func NewListHandler(s ListsService) ListHandler {
	return ListHandler{service: s}
}

func (h *ListHandler) create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"create": "res"})
}

func (h *ListHandler) get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"get": "res"})
}

func (h *ListHandler) getAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getAll": "res"})
}
