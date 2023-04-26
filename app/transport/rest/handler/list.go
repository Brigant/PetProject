package handler

import (
	"errors"
	"net/http"

	"github.com/Brigant/PetPorject/app/core"
	"github.com/Brigant/PetPorject/logger"
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

// Handler for the list creatation.
func (h *ListHandler) create(c *gin.Context) {
	var list core.MovieList

	if err := c.ShouldBindJSON(&list); err != nil {
		h.logger.Debugw("bind json error: %w", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctxAccountID, ok := c.Get(userCtx)
	if !ok {
		h.logger.Debugw("get from contex: %w", core.ErrContexAccountNotFound)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": core.ErrContexAccountNotFound.Error()})

		return
	}

	accountID, err := uuid.Parse(ctxAccountID.(string))
	if err != nil {
		h.logger.Debugw("uuid parse error: %w", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	list.AccountID = accountID

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

// Handler for the movie list getting of the athenticated accaount.
// The path may cointain the query parameters. ex.:
// /list/?type=wish&type=favorite .
func (h *ListHandler) getAll(c *gin.Context) {
	accountID, ok := c.Get(userCtx)
	if !ok {
		h.logger.Debugw("getAll hendler", "error", core.ErrContexAccountNotFound)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": core.ErrContexAccountNotFound.Error()})

		return
	}

	filter := []core.QuerySliceElement{
		{Key: "account_id", Val: accountID.(string)},
	}

	for _, elem := range c.QueryArray("type") {
		filter = append(filter, core.QuerySliceElement{
			Key: "type",
			Val: elem,
		})
	}

	movieLists, err := h.service.GetAllAccountLists(filter)
	if err != nil {
		if errors.Is(err, core.ErrUnkownConditionKey) {
			h.logger.Debugw("getAll hendler", "error", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		h.logger.Debugw("getAll hendler", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	if movieLists == nil {
		h.logger.Debugw("getAll result", "alert", core.ErrNotFound.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"alert": core.ErrNotFound.Error()})

		return
	}

	c.JSON(http.StatusOK, movieLists)
}

type requesMovieList struct {
	ListID  uuid.UUID `json:"list_id" db:"list_id" binding:"required"`
	MovieID uuid.UUID `json:"movie_id" db:"movie_id" binding:"required"`
}

func (h ListHandler) movieToList(c *gin.Context) {
	input := requesMovieList{}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Debugw("Handler movieToList -> ShouldBindJSON", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := h.service.AddMovieToList(input.ListID.String(), input.MovieID.String()); err != nil {
		if errors.Is(err, core.ErrDuplicateRow) || errors.Is(err, core.ErrForeignKeyViolation) {
			h.logger.Debugw("Handler movieToList -> AddMovieToList", "error", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		h.logger.Debugw("Handler movieToList -> AddMovieToList", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"action": "succesfful"})
}
