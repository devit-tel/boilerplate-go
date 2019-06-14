package logger

import (
	"github.com/NV4RE/boilerplate-go/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
)

func InitLog(logConfig config.LogConfig) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(logConfig.Level)
	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Fatal("Error during log level parse:", err)
	}
}
