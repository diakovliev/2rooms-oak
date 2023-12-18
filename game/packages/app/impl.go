package app

import (
	"context"

	"github.com/pior/runnable"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type impl struct {
	ctx        context.Context
	logger     zerolog.Logger
	cancel     context.CancelFunc
	manager    runnable.AppManager
	channel    chan error
	shutdowner fx.Shutdowner
}

func newapp(
	pctx context.Context,
	logger zerolog.Logger,
	manager runnable.AppManager,
	shutdowner fx.Shutdowner,
) (ret *impl, cancel context.CancelFunc) {
	ctx, cancel := context.WithCancel(pctx)
	ret = &impl{
		logger:     logger,
		ctx:        ctx,
		cancel:     cancel,
		manager:    manager,
		channel:    make(chan error, 1),
		shutdowner: shutdowner,
	}
	return
}

func (ra impl) start() {
	ra.logger.Info().Msg("start app...")
	defer ra.logger.Info().Msg("app started")
	go func() {
		err := runnable.Signal(ra.manager.Build()).Run(ra.ctx)
		if err != nil && err != context.Canceled {
			ra.logger.Panic().Err(err).Msg("unexpected error")
		}
		ra.channel <- err
		ra.shutdowner.Shutdown(fx.ExitCode(0))
	}()
}

func (ra impl) stop() error {
	ra.logger.Info().Msg("stop app...")
	defer ra.logger.Info().Msg("app stopped")
	ra.cancel()
	return <-ra.channel
}
