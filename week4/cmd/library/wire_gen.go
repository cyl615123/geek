// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/cyl615123/geek/week4/internal/biz"
	"github.com/cyl615123/geek/week4/internal/config"
	"github.com/cyl615123/geek/week4/internal/data"
	"github.com/cyl615123/geek/week4/internal/server"
	"github.com/cyl615123/geek/week4/internal/service"
	"github.com/micro/go-micro"
)

// Injectors from wire.go:

func initApp(conf *config.Conf) (micro.Service, func(), error) {
	dataData, cleanup, err := data.NewData(conf)
	if err != nil {
		return nil, nil, err
	}
	bookRepo := data.NewBookRepo(dataData)
	bookUseCase := biz.NewBookUseCase(bookRepo)
	libraryService := service.NewLibraryService(bookUseCase)
	httpServer := server.NewHTTPServer(conf, libraryService)
	microService := newApp(httpServer)
	return microService, func() {
		cleanup()
	}, nil
}
