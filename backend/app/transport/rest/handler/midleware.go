package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	athoriazahionHeader = "Authorization"
	userCtx             = "userID"
	roleCtx             = "userRole"
	headerPartsNumber   = 2
	roleAdmin           = "admin"
)

var (
	errEmptyHeader   = errors.New("empty header")
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

func (h Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(athoriazahionHeader)
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
	h.userIdentity(c)

	role, exist := c.Get(roleCtx)
	if !exist {
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
