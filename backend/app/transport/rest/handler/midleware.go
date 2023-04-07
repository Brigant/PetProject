package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authoriazahionHeader = "Authorization"
	authorizationType    = "Bearer"
	userCtx              = "userID"
	roleCtx              = "userRole"
	headerPartsNumber    = 2
	roleAdmin            = "admin"
)

var (
	errEmptyHeader   = errors.New("empty header, expecting Authorization header")
	errInvalidHeader = errors.New("invalid header")
	errEmptyRole     = errors.New("empty role")
	errNotAdmin      = errors.New("you are not admin")
)

func (h Handler) midlewareWithLogger(c *gin.Context) {
	h.log.Debugw("Request:",
		"Method", c.Request.Method,
		"URL", c.Request.URL,
		"User-Agent", c.Request.UserAgent(),
	)
	c.Next()
}

// The middleware checks if there is some registred user.
func (h Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authoriazahionHeader)
	if header == "" {
		h.log.Debugw("userIdentify", "error", errEmptyHeader.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errEmptyHeader.Error(),
		})

		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != headerPartsNumber {
		h.log.Debugw("userIdentify", "error", errInvalidHeader.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errInvalidHeader.Error(),
		})

		return
	}

	if headerParts[0] != authorizationType {
		h.log.Debugw("userIdentify", "error", errInvalidHeader.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errInvalidHeader.Error(),
		})

		return
	}

	if headerParts[1] == "" {
		h.log.Debugw("userIdentify", "error", errInvalidHeader.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errInvalidHeader.Error(),
		})

		return
	}

	userID, userRole, err := h.Account.service.ParseToken(headerParts[1])
	if err != nil {
		h.log.Debugw("userIdentify", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Set(userCtx, userID)
	c.Set(roleCtx, userRole)
}

// This midleware implement the functionality of userIdentity
// and also check the admin role is present.
func (h Handler) adminIdentity(c *gin.Context) {
	role, exist := c.Get(roleCtx)
	if !exist || role == "" {
		h.log.Debugw("adminIdentity", "error", errEmptyRole.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errEmptyRole.Error(),
		})

		return
	}

	if role != roleAdmin {
		h.log.Debugw("adminIdentity", "error", errNotAdmin.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errNotAdmin.Error(),
		})

		return
	}
}
