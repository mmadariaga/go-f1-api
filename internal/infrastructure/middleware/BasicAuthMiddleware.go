package middleware

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
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

		username, password, ok := req.BasicAuth()

		if !ok || username != BASIC_AUTH_USERNAME || password != BASIC_AUTH_PASSWORD {

			authErrorMsg := fmt.Sprintf("Invalid username '%v' with password '%v'", username, password)
			log.Info().Msg(authErrorMsg)
			unauthorize(response)
			return
		}

		callNextMiddleware()
	})
}

func unauthorize(response http.ResponseWriter) {
	http.Error(response, "Unauthorized", http.StatusUnauthorized)
}
