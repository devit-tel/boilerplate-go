package main

import (
	"github.com/NV4RE/boilerplate-go/pkg/commons"
	"github.com/NV4RE/boilerplate-go/pkg/server"
)

var options commons.Options

func init() {
	commons.LoadEnv()
	commons.InitLog()
}

func main() {
	server.CreateServer(options)
}
