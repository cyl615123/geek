package main

import (
	"github.com/BurntSushi/toml"
	"github.com/cyl615123/geek/week4/internal/config"
	"github.com/micro/go-micro"
	microHttp "github.com/micro/go-plugins/server/http"
	"net/http"
)

func main() {
	confFile := "./config/config.toml"
	cfg := config.Conf{}
	_, err := toml.DecodeFile(confFile, &cfg)
	panic(err)

	app, cleanup, err := initApp(&cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(hs *http.Server) micro.Service {
	srv := microHttp.NewServer()
	err := srv.Handle(srv.NewHandler(hs.Handler))
	if err != nil {
		panic(err)
	}
	return micro.NewService(
		micro.Name("library"),
		micro.Server(srv),
	)
}
