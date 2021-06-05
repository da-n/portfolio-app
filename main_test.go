package main

import (
	"os"
	"testing"
)

func Test_it_should_return_error_when_no_env_vars_have_been_set(t *testing.T) {
	// given, when
	err := envCheck()

	// then
	if err == nil {
		t.Errorf("env check should have failed, instead received nil")
	}
}

func Test_it_should_return_nil_when_env_vars_have_been_set(t *testing.T) {
	// given
	os.Setenv("SERVER_ADDRESS", "localhost")
	os.Setenv("SERVER_PORT", "8080")

	// when
	err := envCheck()

	// then
	if err != nil {
		t.Errorf("env check should have passed, instead received error")
	}
}