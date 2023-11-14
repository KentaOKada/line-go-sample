package line

import (
	"github.com/labstack/echo/v4"
)

func (di DefaultInteractor) SendMessage(ctx echo.Context, message string) error {
	err := di.client.SendMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
