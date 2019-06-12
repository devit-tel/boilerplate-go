// Package commons provides data structures used between packages
package commons

// Options provides overall configuration data
type Options struct {
	ServerPort string
	LogLevel   string
}

type LogConfig struct {
	Level string
}

type ServerConfig struct {
	Port string
}

type MongoDbConfig struct {
	Uri string
}

type Config struct {
	Env     string
	Log     LogConfig
	Server  ServerConfig
	MongoDb MongoDbConfig
}
