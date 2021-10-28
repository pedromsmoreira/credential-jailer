package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVaultReader(t *testing.T) {
	t.Run("read existing credential from source should return unencoded credential", func(t *testing.T) {
		expected := &Credential{
			Name:  "name",
			Value: "pwd",
		}
		reader := &EnvVarReader{
			proxy: &MockSystemProxy{
				mValue: "pwd",
				mError: nil},
		}

		cred, err := reader.Read("name")

		assert.NotNil(t, cred)
		assert.Nil(t, err)
		assert.Equal(t, expected.Name, cred.Name)
		assert.Equal(t, expected.Value, cred.Value)
	})

	t.Run("read non existing credential from source should return error", func(t *testing.T) {
		reader := &EnvVarReader{
			proxy: &MockSystemProxy{
				mValue: "",
				mError: errors.New("error fetching env var")},
		}

		cred, err := reader.Read("name")

		assert.NotNil(t, err)
		assert.Nil(t, cred)
	})

	t.Run("read credential from source cannot access system should return error", func(t *testing.T) {
		reader := &EnvVarReader{
			proxy: &MockSystemProxy{
				mValue: "",
				mError: errors.New("cannot access system env vars")},
		}

		cred, err := reader.Read("name")

		assert.NotNil(t, err)
		assert.Nil(t, cred)
	})
}
