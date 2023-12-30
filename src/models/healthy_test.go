package models

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewHealthy(t *testing.T) {
	os.Setenv("DOMAIN", "test.com")
	healthy := NewHealthy()
	assert.NotNil(t, healthy)
}
