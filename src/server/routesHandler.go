package server

import (
	"encoding/json"
	"morsify/src/morse"
	"morsify/src/utils"
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
	input := r.URL.Query().Get("text")
	status, prompt := utils.VerifyPrompt(input)
	status, result := morse.TextToMorse(status, prompt)

	if result == "" || result == " " {
		status = "fail : invalid character detected"
	} else {
		result = result[1:]
	}
	if input == "" || input == " " {
		status, result = "fail : input is empty", ""
	}

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
	input := r.URL.Query().Get("morse")
	status, result := morse.MorseToText(input)

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
