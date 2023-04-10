package handler

import (
	"errors"
	"net/http"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieHandler struct {
	service MovieService
	logger  *logger.Logger
}

func NewMovieHandler(s MovieService, log *logger.Logger) MovieHandler {
	return MovieHandler{
		service: s,
		logger:  log,
	}
}

// Handler for the movie creation.
func (h *MovieHandler) create(c *gin.Context) {
	var movie core.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		h.logger.Debugw("Should bind with movie enteties", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	_, err := uuid.Parse(movie.DirectorID)
	if err != nil {
		h.logger.Debugw("Parse: directorID not uuid", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	_, err = h.service.CreateMovie(movie)
	if err != nil {
		if errors.Is(err, core.ErrForeignViolation) {
			h.logger.Debugw("CreateMovie", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		if errors.Is(err, core.ErrUniqueMovie) {
			h.logger.Debugw("CreateMovie", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		h.logger.Errorw("CreateMovie", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	// c.JSON(http.StatusCreated, gin.H{"action": "successful"})
	c.JSON(http.StatusCreated, movie)
}

func (h *MovieHandler) get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"get": "res"})
}

func (h *MovieHandler) getAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getAll": "res"})
}
