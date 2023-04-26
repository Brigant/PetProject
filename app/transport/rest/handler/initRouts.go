package handler

import (
	"errors"

	"github.com/Brigant/PetPorject/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	availableRoles                  = []string{"user", "admin"}
	errValidatorBind                = errors.New("can't bind the validator")
	checkRoleFunc    validator.Func = func(fl validator.FieldLevel) bool {
		role, ok := fl.Field().Interface().(string)
		if ok {
			for _, r := range availableRoles {
				if r == role {
					return true
				}
			}
		}

		return false
	}
)

// The structure describes the dependencies.
type Deps struct {
	AccountService  AccountService
	DirectorService DirectorService
	MovieService    MovieService
	ListService     ListsService
}

type Handler struct {
	Account  AccountHandler
	Director DirectorHandler
	Movie    MovieHandler
	List     ListHandler
	log      *logger.Logger
}

func NewHandler(deps Deps, logger *logger.Logger) Handler {
	return Handler{
		Account:  NewAccountHandler(deps.AccountService, logger),
		Director: NewDirectorHandler(deps.DirectorService, logger),
		Movie:    NewMovieHandler(deps.MovieService, logger),
		List:     NewListHandler(deps.ListService, logger),
		log:      logger,
	}
}

func (h *Handler) InitRouter(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("checkRole", checkRoleFunc); err != nil {
			h.log.Errorw("bind validator", "err", errValidatorBind.Error())
		}
	}

	router.Use(gin.Recovery(), h.midlewareWithLogger)

	auth := router.Group("/auth")
	{
		auth.POST("/", h.Account.singUp)
		auth.POST("/login", h.Account.login)
		auth.GET("/logout", h.userIdentity, h.Account.logout)
		auth.POST("/refresh", h.Account.refreshToken)
	}

	director := router.Group("/director", h.userIdentity)
	{
		director.POST("/", h.adminIdentity, h.Director.create)
		director.GET("/:id", h.Director.get)
		director.GET("/all", h.Director.getAll)
	}

	movie := router.Group("/movie", h.userIdentity)
	{
		movie.POST("/", h.adminIdentity, h.Movie.create)
		movie.GET("/:id", h.Movie.get)
		movie.GET("/", h.Movie.getAll)
	}

	list := router.Group("list", h.userIdentity)
	{
		list.POST("/", h.List.create)
		list.GET("/:id", h.List.get)
		list.GET("/", h.List.getAll)
		list.POST("/add", h.List.movieToList)
	}

	return router
}
