package config

import (
	"log"
	"os"
	"strconv"
)

func GetApplicationPort() int {
	port, err := strconv.ParseInt(getEnvironmentValue("APP_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("failed to parse APP_PORT: %v", err)
	}

	return int(port)
}

func getEnvironmentValue(key string) string {
	var val string
	if val = os.Getenv(key); val == "" {
		log.Fatalf("environment variable %s is not set", key)
	}

	return val
}
