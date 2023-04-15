package handler

import (
	"errors"
	"fmt"
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

	if err := h.service.CreateMovie(movie); err != nil {
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

	c.JSON(http.StatusCreated, gin.H{"action": "successful"})
}

// Handler for the movie receiving.
func (h *MovieHandler) get(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		h.logger.Debugw("No movieID in the path")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No movieID param in path"})

		return
	}

	_, err := uuid.Parse(id)
	if err != nil {
		h.logger.Debugw("ID is not UUID", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	movie, err := h.service.Get(id)
	if err != nil {
		if errors.Is(err, core.ErrMovieNotFound) {
			h.logger.Debugw("Get movie", "error", err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

			return
		}

		h.logger.Errorw("Get movie", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, movie)
}

// Handler is for the movie's list recievcing weighted by parameters. The full example of url query:
// /?offset=3&filter[genre]=adventure&filter[rate]=10&s[duration]=desc&s[rate]=asc&s[release_date]=asc&limit=100&export=csv
// The allowed values for s[...] are "desc" or "asc".
func (h *MovieHandler) getAll(c *gin.Context) {
	var qp core.QueryParams

	qp.Limit = c.Query("limit")
	qp.Offset = c.Query("offset")
	qp.Filter = c.QueryMap("f")
	qp.Sort = c.QueryMap("s")
	qp.Export = c.Query("export")

	qp.SetDefaultValues()

	if err := qp.Validate(); err != nil {
		h.logger.Debugw("validation", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	movieList, err := h.service.GetList(qp)
	if err != nil {
		if errors.Is(err, core.ErrMovieNotFound) {
			h.logger.Debugw("bad query", "alert", err.Error())
			c.JSON(http.StatusOK, gin.H{"alert": err.Error()})

			return
		}

		h.logger.Debugw("Service Getlist", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	fmt.Printf("%+v\n\n", qp)

	c.JSON(http.StatusOK, movieList)
}
