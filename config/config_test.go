package config_test

import (
	"testing"

	"github.com/MatheusAbdias/simple_payment_service/config"
	"github.com/stretchr/testify/assert"
)

func TestShouldBeLoadConfig(t *testing.T) {
	config, err := config.LoadConfig("../")

	assert.NoError(t, err)
	assert.NotEmpty(t, config.Driver)
	assert.NotEmpty(t, config.Source)
	assert.NotEmpty(t, config.Port)
}
