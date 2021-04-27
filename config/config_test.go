package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	config, err := LoadConfig("..")

	assert.Nil(t, err)
	assert.NotNil(t, config)

	assert.Equal(t, "3001", config.Port)
	assert.Equal(t, "./certs/localhost.crt", config.CertFileLocation)
	assert.Equal(t, "./certs/localhost.key", config.KeyFileLocation)
}
