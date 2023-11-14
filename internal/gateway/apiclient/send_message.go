package apiclient

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

func (r Repository) SendMessage(ctx echo.Context, message string) error {
	// アクセストークンはここに記載
	accessToken := ""

	URL := "https://notify-api.line.me/api/notify"

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	c := &http.Client{}

	form := url.Values{}
	form.Add("message", message)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}
