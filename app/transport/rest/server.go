package rest

import (
	"fmt"
	"net/http"

	"github.com/Brigant/PetPorject/app/repositorie/pg"
	"github.com/Brigant/PetPorject/app/service"
	"github.com/Brigant/PetPorject/app/transport/rest/handler"
	"github.com/Brigant/PetPorject/config"
	"github.com/Brigant/PetPorject/logger"
	_ "github.com/lib/pq" // the blank import is needed beceause of sqlx requirements
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, router http.Handler) *Server {
	server := new(Server)

	server.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return server
}

func SetupAndRun() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return fmt.Errorf("cannot read config: %w", err)
	}

	logger, err := logger.New(cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("cannot create logger: %w", err)
	}

	defer logger.Flush()

	db, err := pg.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("error while creating connection to database: %w", err)
	}

	storage := pg.NewRepository(db)

	services := service.New(
		service.Deps{
			AccountStorage:  storage.AccountDB,
			DirectorStorage: storage.DirectorDB,
			MovieStorage:    storage.MovieDB,
			ListSorage:      storage.ListDB,
		}, cfg)

	restHandlers := handler.NewHandler(
		handler.Deps{
			DirectorService: services.Director,
			AccountService:  services.Account,
			MovieService:    services.Movie,
			ListService:     services.List,
		}, logger)

	routes := restHandlers.InitRouter(cfg.Server.Mode)

	server := NewServer(cfg.Server.Port, routes)

	if err := server.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("cannot run server: %w", err)
	}

	return nil
}
