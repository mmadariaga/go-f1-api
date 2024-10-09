package application_pong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPong(t *testing.T) {

	assert := assert.New(t)

	resp := Pong()

	assert.Equal(resp, "pong")
}
