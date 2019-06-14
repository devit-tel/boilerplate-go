package main

import (
	"github.com/NV4RE/boilerplate-go/pkg/config"
	"github.com/NV4RE/boilerplate-go/pkg/http/server"
	"github.com/NV4RE/boilerplate-go/pkg/logger"
)

func init() {
}

func main() {
	cfg := config.GetConfig()
	logger.InitLog(cfg.Log)
	sv := server.CreateServer(cfg.Server.Addr)
	sv.Listen()
}
