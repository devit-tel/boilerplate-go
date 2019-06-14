package main

import (
    "github.com/NV4RE/boilerplate-go/pkg/logger"
    "github.com/NV4RE/boilerplate-go/pkg/http/server"
    "github.com/NV4RE/boilerplate-go/pkg/config"
)

func init() {
}

func main() {
    cfg := config.GetConfig()
    logger.InitLog(cfg.Log)
	server.CreateServer(cfg.Server.Addr)
}
