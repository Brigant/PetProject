package handler

import (
	"errors"
	"net/http"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

type inputAccountData struct {
	Phone    string `binding:"required,e164,lowercase"`
	Password string `binding:"required,min=8,max=255,ascii"`
}

type inputRefreshToken struct {
	//nolint:tagliatelle
	RefreshToken uuid.UUID `json:"refreshToken"`
}

var errInvalidRefreshToken = errors.New("invalid refresh token")

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
	session := core.Session{
		RequestHost: c.Request.Host,
		ClientIP:    c.ClientIP(),
		UserAgent:   c.Request.UserAgent(),
	}

	var accInputData inputAccountData

	if err := c.ShouldBindJSON(&accInputData); err != nil {
		h.logger.Debugw("ShouldBindJSON", "err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	tokenPair, err := h.service.Login(accInputData.Phone, accInputData.Password, session)
	if err != nil {
		if errors.Is(err, core.ErrUserNotFound) {
			h.logger.Debugw("Login", "alert", err.Error())
			c.JSON(http.StatusNotFound, err.Error())

			return
		}

		h.logger.Errorw("Login", "error", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, tokenPair)
}

func (h AccountHandler) refreshToken(c *gin.Context) {
	var inputToken inputRefreshToken

	if err := c.ShouldBindJSON(&inputToken); err != nil {
		h.logger.Debugw("ShouldBindJSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var badToken inputRefreshToken

	if inputToken == badToken {
		h.logger.Debugw("ShouldBindJSON", "error", errInvalidRefreshToken.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidRefreshToken.Error()})

		return
	}

	session := core.Session{
		RefreshToken: inputToken.RefreshToken.String(),
		RequestHost:  c.Request.Host,
		ClientIP:     c.ClientIP(),
		UserAgent:    c.Request.UserAgent(),
	}

	tokenPair, err := h.service.RefreshTokenpair(session)
	if err != nil {
		h.logger.Errorw("error happened while RefreshTokenpair", "error", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, tokenPair)
}

func (h AccountHandler) logout(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		h.logger.Errorw("logout", "error", core.ErrNotAuthenticated.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": core.ErrNotAuthenticated.Error()})

		return
	}

	accountID, isString := userID.(string)
	if !isString {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accountID is not string"})

		return
	}

	if err := h.service.Logout(accountID); err != nil {
		h.logger.Errorw("logout", "error", err.Error())

		if errors.Is(err, core.ErrNoRowsEffected) {
			c.JSON(http.StatusAccepted, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"action": "successful"})
}
