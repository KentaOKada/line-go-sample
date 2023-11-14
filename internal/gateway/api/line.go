package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (svr *LineServiceServer) SendMessage(c echo.Context) error {
	param := new(SendMessageParam)
	if err := c.Bind(param); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(param.Message)
	err := svr.interactor.SendMessage(c, param.Message)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
