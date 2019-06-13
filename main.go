package main

import (
	"github.com/NV4RE/boilerplate-go/pkg/commons"
	"github.com/NV4RE/boilerplate-go/pkg/server"
	"github.com/Sirupsen/logrus"
)

var options commons.Options

func init() {
	commons.InitLog()
	logrus.Info(commons.GetConfig())
}

func main() {
	server.CreateServer(options)
}
