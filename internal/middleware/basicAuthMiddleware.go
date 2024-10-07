package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const (
	BASIC_AUTH_USERNAME = "go"
	BASIC_AUTH_PASSWORD = "api"
)

// Require basic auth everywhere but /ping
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, req *http.Request) {

		callNextMiddleware := func() {
			next.ServeHTTP(response, req)
		}

		path := req.URL.Path

		if path == "/ping" {
			callNextMiddleware()
			return
		}

		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			unauthorize(response)
			return
		}

		if !strings.HasPrefix(authHeader, "Basic ") {
			unauthorize(response)
			return
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			unauthorize(response)
			return
		}

		credentials := strings.SplitN(string(decodedCredentials), ":", 2)
		if len(credentials) != 2 {
			unauthorize(response)
			return
		}

		username, password := credentials[0], credentials[1]

		if username != BASIC_AUTH_USERNAME || password != BASIC_AUTH_PASSWORD {
			unauthorize(response)
			return
		}

		callNextMiddleware()
	})
}

func unauthorize(response http.ResponseWriter) {
	http.Error(response, "Unauthorized", http.StatusUnauthorized)
}
