package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(s AccountService) AccountHandler {
	return AccountHandler{service: s}
}

func (h *AccountHandler) create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"create": "res"})
}

func (h *AccountHandler) get(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"get": "res"})
}
