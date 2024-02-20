package server

import (
	"github.com/abulo/ratel/v3/core/app"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/sirupsen/logrus"
)

type Engine struct {
	app.Application
}

// NewHertzEngine Hertz框架
func NewHertzEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.NewHertzServer,
	); err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panic("startup")
	}
	return eng
}

// NewGrpcEngine Grpc框架
func NewGrpcEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.NewGrpcServer,
	); err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panic("startup")
	}
	return eng
}
