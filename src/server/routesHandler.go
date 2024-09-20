package server

import (
	"encoding/json"
	"morsify/src/morse"
	"net/http"
)

func VerifyMethod(r *http.Request, method string) bool {
	return r.Method == method
}

func (s *Server) MosrsifyHandler(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, http.MethodGet) {
		http.Error(w, "Error: Invalid method", http.StatusMethodNotAllowed)
		return
	}
	prompt := r.URL.Query().Get("text")

	status, result := morse.TextToMorse(prompt)

	response := struct {
		Status string `json:"status"`
		Result string `json:"result"`
	}{
		Status: status,
		Result: result,
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) DemorsifyHandler(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, http.MethodGet) {
		http.Error(w, "Error: Invalid method", http.StatusMethodNotAllowed)
		return
	}

}
