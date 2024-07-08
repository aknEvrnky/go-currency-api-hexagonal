package config

import (
	"errors"
	"os"
	"strconv"
)

var (
	ErrEnvKeyIsNotSet       = errors.New("environment variable is not set")
	ErrEnvValueTypeMismatch = errors.New("environment variable type mismatch")
)

func GetApplicationPort() int {
	port, err := strconv.ParseInt(getEnvironmentValue("APP_PORT"), 10, 64)
	if err != nil {
		panic(ErrEnvValueTypeMismatch)
	}

	return int(port)
}

func getEnvironmentValue(key string) string {
	var val string
	if val = os.Getenv(key); val == "" {
		panic(ErrEnvKeyIsNotSet)
	}

	return val
}
