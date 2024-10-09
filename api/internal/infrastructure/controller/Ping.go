package infrastructure_controller

import (
	"net/http"

	application_pong "github.com/mmadariaga/go-api/internal/application/pong"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := application_pong.Pong()
	w.Write([]byte(resp))
}
