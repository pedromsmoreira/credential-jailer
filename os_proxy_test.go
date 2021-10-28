package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOsProxy(t *testing.T) {
	t.Run("read existing env var should return value", func(t *testing.T) {
		// Arrange
		os.Setenv("test", "var-exists")
		proxy := &OsProxy{}
		// Act
		v, err := proxy.ReadEnvVar("test")
		// Assert
		assert.Equal(t, "var-exists", v)
		assert.Nil(t, err)
		// Clean Up
		os.Unsetenv("test")
	})

	t.Run("Read env var does not exist should return empty value and error", func(t *testing.T) {
		// Arrange
		proxy := &OsProxy{}
		// Act
		v, err := proxy.ReadEnvVar("not exists")
		// Assert
		assert.Equal(t, "", v)
		assert.NotNil(t, err)
	})
}
