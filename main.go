package main

import (
	"github.com/devit-tel/boilerplate-go/pkg/config"
	"github.com/devit-tel/boilerplate-go/pkg/http/server"
	"github.com/devit-tel/boilerplate-go/pkg/logger"
)

func init() {
}

func main() {
	cfg := config.GetConfig()
	logger.InitLog(cfg.Log)
	sv := server.CreateServer(cfg.Server.Addr)
	sv.Listen()
}
