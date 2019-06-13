package commons

import (
	"github.com/Sirupsen/logrus"
	"os"
)

func InitLog() {
	env := GetConfig()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	level, err := logrus.ParseLevel(env.Log.Level)
	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Fatal("Error during log level parse:", err)
	}
}
