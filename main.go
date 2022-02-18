package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/io4io/terra/api"
	"github.com/io4io/terra/internal"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	internal.ConfigureLogger()
	defer zap.L().Sync()

	// Init modules.

	// Run the server.
	addr := ":8080"
	httpServer := &http.Server{
		Addr:    addr,
		Handler: api.Router(),
	}

	ctx, cancel := context.WithCancel(context.Background()) // for gracefully shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan,
			syscall.SIGHUP,  // kill -SIGHUP (hang up, 1) xxxx
			syscall.SIGINT,  // kill -SIGINT (interrupt, 2) xxxx, Ctrl+c
			syscall.SIGTERM, // kill -SIGTERM (software termination, 15), kill a program gracefully
		)
		<-sigChan
		cancel()
	}()

	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gctx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		zap.S().Warnw("program exited", zap.Error(err))
	}
}
