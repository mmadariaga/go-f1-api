package controller

import (
	"net/http"

	application "github.com/mmadariaga/go-api/internal/application/pong"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := application.Pong()
	w.Write([]byte(resp))
}
