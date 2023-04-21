package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
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
		if errors.Is(err, core.ErrNotFound) {
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
// /movie/?offset=3&f=genre:comedy&f=rate:10&s=duration:desc&s=rate:asc&s=release_date:asc&limit=100&export=csv
// The allowed values for s[...] are "desc" or "asc", for export: "csv" or "none".
func (h *MovieHandler) getAll(c *gin.Context) {
	queryParameter, err := h.prepareQueryParams(c)
	if err != nil {
		h.logger.Debugw("prepareQueryParams", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	switch queryParameter.Export {
	case "csv":
		movieList, err := h.service.GetCSV(queryParameter)
		if err != nil {
			if errors.Is(err, core.ErrNotFound) {
				h.logger.Debugw("bad query", "alert", err.Error())
				c.JSON(http.StatusOK, gin.H{"alert": err.Error()})

				return
			}

			h.logger.Debugw("Service Getlist", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		csvList, err := gocsv.MarshalBytes(movieList)
		if err != nil {
			h.logger.Debugw("Marshal CSV", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.Data(http.StatusOK, "text/csv; charset=utf-8", csvList)

	default:
		movieList, err := h.service.GetList(queryParameter)
		if err != nil {
			if errors.Is(err, core.ErrNotFound) {
				h.logger.Debugw("bad query", "alert", err.Error())
				c.JSON(http.StatusOK, gin.H{"alert": err.Error()})

				return
			}

			h.logger.Debugw("Service Getlist", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, movieList)
	}
}

func (h MovieHandler) prepareQueryParams(c *gin.Context) (core.ConditionParams, error) {
	var queryParameter core.ConditionParams
	queryParameter.CheckList.Export = true
	queryParameter.CheckList.Filter = true
	queryParameter.CheckList.Sort = true
	queryParameter.CheckList.Offset = true
	queryParameter.CheckList.Limit = true

	queryParameter.Limit = c.Query("limit")
	queryParameter.Offset = c.Query("offset")
	queryParameter.Export = c.Query("export")

	for _, v := range c.QueryArray("f") {
		keyval := strings.Split(v, ":")

		var element core.QuerySliceElement

		element.Key = keyval[0]
		element.Val = keyval[1]

		queryParameter.Filter = append(queryParameter.Filter, element)
	}

	for _, v := range c.QueryArray("s") {
		keyval := strings.Split(v, ":")

		var element core.QuerySliceElement

		element.Key = keyval[0]
		element.Val = keyval[1]

		queryParameter.Sort = append(queryParameter.Sort, element)
	}

	queryParameter.SetDefaultValues()

	if err := queryParameter.Validate(); err != nil {
		return core.ConditionParams{}, fmt.Errorf("query preparetion failed: %w", err)
	}

	return queryParameter, nil
}
