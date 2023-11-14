package app

import (
	"context"

	"github.com/KentaOKada/line-go-sample/internal/config"
	"github.com/KentaOKada/line-go-sample/internal/domain/usecase/line"
	"github.com/KentaOKada/line-go-sample/internal/gateway/api"
	clientrepo "github.com/KentaOKada/line-go-sample/internal/gateway/apiclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(conf *config.Config) (err error) {
	initializer := &ServerInitializer{}
	return startProcess(conf, initializer)
}

type ServerInitializer struct{}

type Server struct {
	e *echo.Echo
}

//nolint:funlen
func (i *ServerInitializer) New(ctx context.Context, conf *config.Config) (Process, error) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client, err := clientrepo.NewRepository()
	if err != nil {
		return nil, err
	}

	lineInteractor := line.NewInteractor(conf, client)

	lineServer := api.NewLineServiceServer(lineInteractor)

	e.POST("", lineServer.SendMessage)

	return &Server{e: e}, nil
}

func (s *Server) Start(ctx context.Context) error {
	return s.e.Start(":1323")
}
