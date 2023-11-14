package main

import (
	"github.com/KentaOKada/line-go-sample/internal/app"
	"github.com/KentaOKada/line-go-sample/internal/config"
)

func main() {
	conf, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}

	if err := app.StartServer(conf); err != nil {
		panic(err)
	}
}
