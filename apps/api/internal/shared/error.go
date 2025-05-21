package shared

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondError(w http.ResponseWriter, code int, message string) {
	log.Println("Error:", message)
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(ErrorResponse{Error: message})
	if err != nil {
		return
	}
}
