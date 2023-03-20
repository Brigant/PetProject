package rest

import (
	"fmt"
	"net/http"

	"github.com/Brigant/PetPorject/backend/app/repositorie/pg"
	"github.com/Brigant/PetPorject/backend/app/service"
	"github.com/Brigant/PetPorject/backend/app/transport/rest/handler"
	"github.com/Brigant/PetPorject/backend/config"
	"github.com/Brigant/PetPorject/backend/logger"
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
		})

	restHandlers := handler.New(
		handler.Deps{
			DirectorService: services.Director,
			AccountService:  services.Account,
		}, logger)

	routes := restHandlers.InitRouter(cfg.Server.Mode)

	server := NewServer(cfg.Server.Port, routes)

	if err := server.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("cannot run server: %w", err)
	}

	return nil
}
