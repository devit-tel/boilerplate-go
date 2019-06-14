package logger

import (
    "github.com/NV4RE/boilerplate-go/pkg/config"
    "github.com/Sirupsen/logrus"
    "os"
)

func InitLog(logConfig config.LogConfig) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	level, err := logrus.ParseLevel(logConfig.Level)
	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Fatal("Error during log level parse:", err)
	}
}
