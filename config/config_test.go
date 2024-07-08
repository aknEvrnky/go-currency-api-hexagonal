package config

import (
	"os"
	"testing"
)

func TestItCanGetEnvValue(t *testing.T) {
	os.Setenv("FOO", "bar")

	if getEnvironmentValue("FOO") != "bar" {
		t.Error("Expected bar, got", getEnvironmentValue("FOO"))
	}

	os.Unsetenv("FOO")

}

func TestItCanGetAppPort(t *testing.T) {
	os.Setenv("APP_PORT", "3000")

	if GetApplicationPort() != 3000 {
		t.Error("Expected 3000, got", GetApplicationPort())
	}

	os.Unsetenv("APP_PORT")
}
