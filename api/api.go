package api

// import (
// 	"encoding/json"
// 	"net/http"
// )

type ReqModel struct {
	Username string
}

type ResModel struct {
	Username string
}

type ApiError struct {
	msg string
}

func (e *ApiError) Error() string {
	return e.msg
}
