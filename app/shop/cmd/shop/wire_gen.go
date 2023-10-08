// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kx-boutique/app/shop/internal/biz"
	"kx-boutique/app/shop/internal/conf"
	"kx-boutique/app/shop/internal/data"
	"kx-boutique/app/shop/internal/server"
	"kx-boutique/app/shop/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	healthService := service.NewHealthService()
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	productRepo := data.NewProductRepo(dataData, logger)
	productUsecase := biz.NewProductUsecase(productRepo, logger)
	productsService := service.NewProductsService(productUsecase)
	grpcServer := server.NewGRPCServer(confServer, healthService, productsService, logger)
	httpServer := server.NewHTTPServer(confServer, healthService, productsService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}