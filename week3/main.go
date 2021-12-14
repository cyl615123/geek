package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	httpServer := http.Server{
		Addr: "127.0.0.1:8080",
	}
	g.Go(func() error {
		err := httpServer.ListenAndServe()
		fmt.Printf("httpServer over. time:%v, err:%v\n", time.Now(), err)
		return err
	})
	g.Go(func() error {
		timer := time.NewTimer(time.Second * 5)
		select {
		case <-ctx.Done():
			fmt.Printf("timer ctx over. time:%v, err:%v\n", time.Now(), ctx.Err())
		case t := <-timer.C:
			fmt.Printf("timer over. time:%v\n", t)
		}
		err := httpServer.Shutdown(context.Background())
		fmt.Printf("httpCloser over. time:%v, err:%v\n", time.Now(), err)
		return err
	})
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Printf("signal ctx over. time:%v, err:%v\n", time.Now(), err)
			return err
		case sig := <-quit:
			err := errors.Errorf("os signal: %v", sig)
			fmt.Printf("signal over. time:%v, err:%v\n", time.Now(), err)
			return err
		}
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("wait over: %v\n", err)
	}
	return
}
