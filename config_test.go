package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := Config{}
	assert.Equal(cfg.data, "")
	assert.Equal(cfg.brand, "")
}

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := Config{data: "/data", brand: "Test"}
	assert.Equal(cfg.data, "/data")
	assert.Equal(cfg.brand, "Test")
}
