package nexrev_utils

import (
	"net/http"
)

func ErrorHTTP(w *http.ResponseWriter, code int, message string) {
	(*w).WriteHeader(code)
	(*w).Header().Set("Content-Type", "text/plain")
	(*w).Header().Set("X-Content-Type-Options", "nosniff")
	(*w).Write([]byte(message))
	println(message)
}
