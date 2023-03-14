package handler

import (
	"github.com/Brigant/GoPetPorject/backend/logger"
	"github.com/gin-gonic/gin"
)

type Deps struct {
	AccountService  AccountService
	DirectorService DirectorService
	MovieService    MovieService
	ListsService    ListsService
}

type Handler struct {
	Account  AccountHandler
	Director DirectorHandler
	Movie    MovieHandler
	List     ListHandler
	log      *logger.Logger
}

func New(deps Deps, logger *logger.Logger) Handler {
	return Handler{
		Account:  NewAccountHandler(deps.AccountService, logger),
		Director: NewDirectorHandler(deps.DirectorService),
		Movie:    NewMovieHandler(deps.MovieService),
		List:     NewListHandler(deps.ListsService),
		log: logger,
	}
}

func (h *Handler) InitRouter(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Recovery(), h.midlewareWithLogger)

	auth := router.Group("/auth")
	{
		auth.POST("/", h.Account.singUp)
		auth.POST("/login", h.Account.login)
		auth.GET("/logout", h.Account.logout)
		auth.POST("/refress", h.Account.refreshToken)
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
