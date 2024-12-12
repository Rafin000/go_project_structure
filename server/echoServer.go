package server

import (
 "fmt"

 cockroachHandlers "go_project_structure/cockroach/handlers"
 cockroachRepositories "go_project_structure/cockroach/repositories"
 cockroachService "go_project_structure/cockroach/services"


 "go_project_structure/config"
 "go_project_structure/database"
 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"
 "github.com/labstack/gommon/log"
)

type echoServer struct {
  app  *echo.Echo
  db   database.Database
  conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
  echoApp := echo.New()
  echoApp.Logger.SetLevel(log.DEBUG)

  return &echoServer{
    app:  echoApp,
    db:   db,
    conf: conf,
  }
}

func (s *echoServer) Start() {
  s.app.Use(middleware.Recover())
  s.app.Use(middleware.Logger())

  // Health check adding
  s.app.GET("v1/health", func(c echo.Context) error {
    return c.String(200, "OK")
  })

  s.initializeCockroachHttpHandler()
  
  serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
  s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCockroachHttpHandler() {
	// Initialize all layers
	cockroachPostgresRepository := cockroachRepositories.NewCockroachPostgresRepository(s.db)
	cockroachFCMMessaging := cockroachRepositories.NewCockroachFCMMessaging()
	
	cockroachService := cockroachService.NewCockroachServiceImpl(
	  cockroachPostgresRepository,
	  cockroachFCMMessaging,
	)
	
	cockroachHttpHandler := cockroachHandlers.NewCockroachHttpHandler(cockroachService)
	
	// Routers
	cockroachRouters := s.app.Group("v1/cockroach")
	cockroachRouters.POST("", cockroachHttpHandler.DetectCockroach)
  }