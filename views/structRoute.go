package views

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response is exported
type Response struct {
	When time.Time `json:"time"`
}

// StructRoute is exported
func StructRoute(w http.ResponseWriter, r *http.Request) {
	t := Response{When: time.Now()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)

}
