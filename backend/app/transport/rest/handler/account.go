package handler

import (
	"errors"
	"net/http"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service AccountService
	logger  *logger.Logger
}

func NewAccountHandler(s AccountService, logger *logger.Logger) AccountHandler {
	return AccountHandler{
		service: s,
		logger:  logger,
	}
}

type signInInput struct {
	Phone    string
	Password string
}

type inputRefreshToken struct {
	RefreshToken string
}

func (h AccountHandler) singUp(c *gin.Context) {
	var account core.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		h.logger.Debugw("ShouldBindJSON", "err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	h.logger.Debugw("signUp", "phone", account.Phone, "age", account.Age)

	userID, err := h.service.CreateUser(account)
	if err != nil {
		if errors.Is(err, core.ErrDuplicatePhone) {
			h.logger.Debugw("CreateUser", "error", err.Error())

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		h.logger.Errorw("CreateUser", "error", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"userID": userID})
}

func (h AccountHandler) login(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"get": "res"})
}

func (h AccountHandler) refreshToken(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"get": "res"})
}

func (h AccountHandler) logout(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"get": "res"})
}
