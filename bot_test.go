package teles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bot(t *testing.T) {
	// Invalid
	b, err := newBot("invalid token")
	assert.Error(t, err)
	assert.Nil(t, b)

	err = b.sendMessage("message")
	assert.Error(t, err)

	// Valid
	b, err = newBot("valid token")
	assert.NoError(t, err)
	assert.NotNil(t, b)

	err = b.sendMessage("message")
	assert.NoError(t, err)
}
