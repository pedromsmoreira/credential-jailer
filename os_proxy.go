package main

import (
	"errors"
	"os"
)

type OsProxy struct {
}

func (op *OsProxy) ReadEnvVar(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if value == "" && !exists {
		return value, errors.New("environment variable does not exist")
	}

	return value, nil
}
