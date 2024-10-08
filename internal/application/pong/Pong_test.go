package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPong(t *testing.T) {

	resp := Pong()

	assert.Equal(t, resp, "pong")
}
