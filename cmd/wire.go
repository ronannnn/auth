//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ronannnn/auth/internal/apis"

	"github.com/google/wire"
)

func InitHttpServer() (*apis.HttpServer, error) {
	wire.Build(wireSet)
	return &apis.HttpServer{}, nil
}
