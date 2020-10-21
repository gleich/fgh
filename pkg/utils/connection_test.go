package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasInternetConnection(t *testing.T) {
	assert.True(t, HasInternetConnection())
}
