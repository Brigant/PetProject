package handler

import (
	"fmt"
	"net/http"

	"github.com/Brigant/PetPorject/backend/app/core"
	"github.com/Brigant/PetPorject/backend/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location,omitempty"`
}

type DirectorHandler struct {
	service DirectorService
	logger  *logger.Logger
}

func NewDirectorHandler(s DirectorService, log *logger.Logger) DirectorHandler {
	return DirectorHandler{
		service: s,
		logger:  log,
	}
}

// Create the director object.
func (h *DirectorHandler) create(c *gin.Context) {
	var director core.Director

	if err := c.ShouldBindJSON(&director); err != nil {
		h.logger.Debugw("Create director", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := h.service.CreateDirector(director); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"action": "successful"})
}

// Returns the object of the director defined by its ID.
// The ID should be passed through the URI path.
func (h *DirectorHandler) get(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		h.logger.Debugw("No direcrotId param in path")
		c.JSON(http.StatusBadRequest, gin.H{"erro": "No direcrotId param in path"})

		return
	}

	_, err := uuid.Parse(id)
	if err != nil {
		h.logger.Debugw("id si not UUID", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	director, err := h.service.GetDirectorWithID(string(id))
	if err != nil {
		h.logger.Errorw("GetDirectorWithID", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	res := []core.Director{director}

	c.JSON(http.StatusOK, res)
}

// Returns the slice of the directors.
func (h *DirectorHandler) getAll(c *gin.Context) {
	directorsList, _ := h.service.GetDirectorList()

	for _, director := range directorsList {
		fmt.Println(director)
	}

	c.JSON(http.StatusOK, directorsList)
}
