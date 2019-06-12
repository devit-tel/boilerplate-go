package commons

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func InitLog() {
	env := GetConfig()

	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	level, err := log.ParseLevel(env.Log.Level)
	if err == nil {
		log.SetLevel(level)
	} else {
		log.Fatal("Error during log level parse:", err)
	}
}
