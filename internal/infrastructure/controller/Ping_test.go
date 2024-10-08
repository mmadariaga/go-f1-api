package controller

import (
	"net/http"
	"testing"

	"github.com/mmadariaga/go-api/internal/test"
)

func TestPing(t *testing.T) {

	mockWriter := new(test.MockResponseWriter)
	mockWriter.On(
		"Write",
		[]byte("pong"),
	).Return(
		len("pong"),
		nil,
	)

	Ping(mockWriter, new(http.Request))

	mockWriter.AssertExpectations(t)
}
