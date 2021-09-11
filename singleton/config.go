package singleton

import (
	"os"
	"sync"
)

type config struct {
	ServerUrl string
	Port      string
}

var (
	once            sync.Once
	singletonConfig *config
)

func GetConfig() *config {
	once.Do(func() {
		singletonConfig = &config{
			ServerUrl: getEnv("server", "localhost"),
			Port:      getEnv("port", "8080"),
		}
	})
	return singletonConfig
}

func getEnv(key, defaultVal string) string {
	if val, exist := os.LookupEnv(key); exist {
		return val
	}
	return defaultVal
}
