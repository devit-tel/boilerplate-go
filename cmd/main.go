package main

import (
	"flag"
	"github.com/NV4RE/boilerplate-go/pkg/commons"
	"github.com/NV4RE/boilerplate-go/pkg/server"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
)

var options commons.Options

func init() {
	pflag.StringVar(&options.ServerPort, "port", "8080", "Server port for listening REST calls")
	pflag.StringVar(&options.LogLevel, "log-level", "info", "Log level, options are panic, fatal, error, warning, info and debug")
}

func main() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	level, err := logrus.ParseLevel(options.LogLevel)
	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Fatal("Error during log level parse:", err)
	}

	// webserver := server.createServer(options)
	// webserver.Start()

	// <-sigs
	// logrus.Warn("Shutting down...")
	// webserver.Stop()
	server.CreateServer(options)

	sigs := make(chan os.Signal, 1)
	<-sigs
	logrus.Warn("Shutting down...")
	// close(stop)
	// wg.Wait()
}
