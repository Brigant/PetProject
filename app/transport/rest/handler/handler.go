package handler

import (
	"github.com/gin-gonic/gin"
)

type Deps struct {
	AccountService  AccountService
	DirectorService DirectorService
	MovieService    MovieService
}

type Handler struct {
	Account  AccountHandler
	Director DirectorHandler
	Movie    MovieHandler
	List     ListHandler
}

func New(deps Deps) Handler {
	return Handler{
		Account:  NewAccountHandler(deps.AccountService),
		Director: NewDirectorHandler(deps.DirectorService),
	}
}

func (h *Handler) InitRouter(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Recovery(), gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/", h.Account.create)
		auth.GET("/", h.Account.get)
	}

	director := router.Group("/director")
	{
		director.POST("/", h.Director.create)
		director.GET("/", h.Director.get)
	}

	movie := router.Group("/movie")
	{
		movie.POST("/", h.Movie.create)
		movie.GET("/{id}", h.Movie.get)
		movie.GET("/", h.Movie.getAll)

	}

	list := router.Group("list")
	{
		list.POST("/", h.List.create)
		list.GET("/{id}", h.List.get)
		list.GET("/", h.List.getAll)
	}

	return router
}
