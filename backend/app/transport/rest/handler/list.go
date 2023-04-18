package handler

import (
	"errors"
	"net/http"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListHandler struct {
	service ListsService
	logger  *logger.Logger
}

func NewListHandler(s ListsService, log *logger.Logger) ListHandler {
	return ListHandler{
		service: s,
		logger:  log,
	}
}

func (h *ListHandler) create(c *gin.Context) {
	var list core.MovieList

	if err := c.ShouldBindJSON(&list); err != nil {
		h.logger.Debugw("bind json error: %w", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ID, ok := c.Get(userCtx)
	if !ok {
		h.logger.Debugw("get from contex: %w", core.ErrContexAccountNotFound)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": core.ErrContexAccountNotFound.Error()})

		return
	}

	accountID, err := uuid.Parse(ID.(string))
	if err != nil {
		h.logger.Debugw("uuid parse error: %w", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	list.AccountID = accountID

	if err := list.Validate(); err != nil {
		h.logger.Debugw("list validation: %w", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	listID, err := h.service.Create(list)
	if err != nil {
		if errors.Is(err, core.ErrDuplicateRow) {
			h.logger.Debugw("service create got an error: %w", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		h.logger.Debugw("service create got an error: %w", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"created with ID": listID})
}

func (h *ListHandler) get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"get": "res"})
}

func (h *ListHandler) getAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getAll": "res"})
}
