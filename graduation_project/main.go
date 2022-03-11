package main

import (
	"context"
	"github.com/cyl615123/geek/graduation_project/controller/download"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	server := http.Server{}
	eg.Go(func() error {
		server.ListenAndServe()
		return nil
	})
	eg.Go(func() error {
		return download.Downloading(ctx)
	})
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			server.Shutdown(ctx)
		}
		return nil
	})
	eg.Wait()
}
