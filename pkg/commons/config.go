package commons

import (
	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

func LoadEnv() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logrus.Info("Not found .env file")
	}
}

func GetConfig() *Config {
	return &Config{
		Env: getEnv("GO_ENV", "LOCAL"),
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "INFO"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", ":8080"),
		},
		MongoDb: MongoDbConfig{
			Uri: getEnv("MONGODB_URI", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logrus.Info("ENV: ", key, " not found use default: ", defaultVal)
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	logrus.Info("ENV: ", key, " not found use default: ", defaultVal)
	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(key string, defaultVal bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	logrus.Info("ENV: ", key, " not found use default: ", defaultVal)
	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(key string, defaultVal []string, sep string) []string {
	valStr := getEnv(key, "")
	if valStr == "" {
		return defaultVal
	}
	val := strings.Split(valStr, sep)
	logrus.Info("ENV: ", key, " not found use default: ", defaultVal)
	return val
}
