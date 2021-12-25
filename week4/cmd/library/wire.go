package main

import (
	"github.com/cyl615123/geek/week4/internal/biz"
	"github.com/cyl615123/geek/week4/internal/config"
	"github.com/cyl615123/geek/week4/internal/data"
	"github.com/cyl615123/geek/week4/internal/server"
	"github.com/cyl615123/geek/week4/internal/service"
	"github.com/google/wire"
	"github.com/micro/go-micro"
)

func initApp(*config.Conf) (micro.Service, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
