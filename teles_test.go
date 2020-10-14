package teles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_teles(t *testing.T) {
	// Invalid
	err := Start("invalid token")
	assert.Error(t, err)
	err = Logger("message")
	assert.Error(t, err)

	// Valid
	err = Start("valid token")
	assert.NoError(t, err)

	err = Logger("message")
	assert.NoError(t, err)
}
