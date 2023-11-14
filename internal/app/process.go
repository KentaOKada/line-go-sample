package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/KentaOKada/line-go-sample/internal/config"
)

type ProcessInitializer interface {
	New(context.Context, *config.Config) (Process, error)
}

type Process interface {
	Start(context.Context) error
}

func startProcess(conf *config.Config, pi ProcessInitializer) (err error) {
	// Context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	defer signal.Stop(interrupt)

	process, err := pi.New(ctx, conf)
	if err != nil {
		return err
	}

	go func() {
		if err := process.Start(ctx); err != nil {
			// TODO: ログ出力
		}

		cancel()
	}()

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	return err
}
