package views

import "net/http"

// ByteRoute is exported
func ByteRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "This was from a byte slice to a string"}`))
}
