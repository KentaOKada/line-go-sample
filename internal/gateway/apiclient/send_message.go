package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (r Repository) SendMessage(ctx echo.Context, message string) error {
	// アクセストークンはここに記載
	accessToken := ""

	// UserIDはここに記載
	userID := ""

	URL := "https://api.line.me/v2/bot/message/push"

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	c := &http.Client{}

	msg := Message{Type: "text", Text: message}
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	jsonStr := `{"to":"` + userID + `",` + `"messages":[` + string(b) + `]}`

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}
