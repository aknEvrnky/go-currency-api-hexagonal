package config

import (
	"github.com/stretchr/testify/assert"
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

func TestItCantGetEnvVariable(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic")
			assert.Error(t, ErrEnvKeyIsNotSet)
		}
	}()

	getEnvironmentValue("FOO")
}

func TestItCantGetEnvVariableTypeMismatch(t *testing.T) {
	os.Setenv("APP_PORT", "foo")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic")
			assert.Error(t, ErrEnvValueTypeMismatch)
		}
	}()

	GetApplicationPort()
}
