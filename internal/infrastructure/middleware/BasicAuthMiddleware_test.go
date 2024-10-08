package infrastructure_middleware

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicAuthMiddleware_PingRequiresNoAuth(t *testing.T) {

	// Expected: 200 Authorized
	req := httpReqFactory(http.MethodGet, "/ping", "")
	recorder := runBasicAuthMiddleware(req)

	respBody := recorder.Body.String()

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Authorized", strings.TrimSpace(respBody))
}

func TestBasicAuthMiddleware_ValidCredentials(t *testing.T) {

	// Expected: 200 Authorized
	req := httpReqFactory(http.MethodGet, "/somepath", BASIC_AUTH_USERNAME)
	recorder := runBasicAuthMiddleware(req)

	respBody := recorder.Body.String()

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Authorized", strings.TrimSpace(respBody))
}

func TestBasicAuthMiddleware_InvalidCredentials(t *testing.T) {

	// Expected: 401 Unauthorized
	req := httpReqFactory(http.MethodGet, "/somepath", "wrongUser")
	recorder := runBasicAuthMiddleware(req)

	respBody := recorder.Body.String()

	require.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "Unauthorized\n", respBody)
}

func TestBasicAuthMiddleware_NoCredentials(t *testing.T) {

	// Expected: 401 Unauthorized
	req := httpReqFactory(http.MethodGet, "/somepath", "")
	recorder := runBasicAuthMiddleware(req)

	respBody := recorder.Body.String()

	require.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "Unauthorized\n", respBody)
}

////////////////////////////////////////
// Helper functions
////////////////////////////////////////

func runBasicAuthMiddleware(req *http.Request) *httptest.ResponseRecorder {

	basicAuthHandler := basicAuthMiddlewareHandlerFactory()
	recorder := httptest.NewRecorder()
	basicAuthHandler.ServeHTTP(recorder, req)

	return recorder
}

func basicAuthMiddlewareHandlerFactory() http.Handler {

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized"))
	})

	handler := BasicAuthMiddleware(next)

	return handler
}

func httpReqFactory(method, path, authUsername string) *http.Request {
	req := httptest.NewRequest(method, path, nil)

	if authUsername != "" {
		req.Header.Set(
			"Authorization",
			createBasicAuthHeader(authUsername, BASIC_AUTH_PASSWORD),
		)
	}

	return req
}

func createBasicAuthHeader(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
