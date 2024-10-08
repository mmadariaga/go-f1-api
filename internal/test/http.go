package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

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
