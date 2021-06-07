package app

import (
	"os"
	"testing"
)

func TestItShouldReturnErrorWhenNoEnvVarsHaveBeenSet(t *testing.T) {
	err := envCheck()

	if err == nil {
		t.Errorf("env check should have failed, instead received nil")
	}
}

func TestItShouldReturnNilWhenEnvVarsHaveBeenSet(t *testing.T) {
	os.Setenv("SERVER_ADDRESS", "localhost")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "badpassword123")
	os.Setenv("DB_ADDRESS", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "portfolio-app")

	err := envCheck()

	if err != nil {
		t.Errorf("env check should have passed, instead received error")
	}
}
