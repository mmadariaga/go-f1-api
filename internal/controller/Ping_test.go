package controller

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestPing(t *testing.T) {

	mockWriter := new(MockResponseWriter)
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

////////////////////////////////////////
// Helper structs & functions
////////////////////////////////////////

// http.ResponseWriter
type MockResponseWriter struct {
	mock.Mock
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}
func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
}
func (m *MockResponseWriter) Header() http.Header {
	args := m.Called()
	return args.Get(0).(http.Header)
}
