package crypt_test

import (
	"testing"

	"github.com/brunomilani/crypt"
	_ "github.com/brunomilani/crypt/apr1_crypt"
	"github.com/stretchr/testify/assert"
)

func TestIsHashSupported(t *testing.T) {
	apr1 := crypt.IsHashSupported("$apr1$salt$hash")
	assert.True(t, apr1)
	other := crypt.IsHashSupported("$unknown$salt$hash")
	assert.False(t, other)
}
