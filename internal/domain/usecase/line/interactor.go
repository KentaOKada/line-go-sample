package line

import (
	"github.com/KentaOKada/line-go-sample/internal/config"
	"github.com/KentaOKada/line-go-sample/internal/gateway/apiclient"
	"github.com/labstack/echo/v4"
)

type Interactor interface {
	SendMessage(ctx echo.Context, message string) error
}

type DefaultInteractor struct {
	conf   *config.Config
	client *apiclient.Repository
}

func NewInteractor(
	conf *config.Config,
	client *apiclient.Repository,
) *DefaultInteractor {
	return &DefaultInteractor{
		conf:   conf,
		client: client,
	}
}
