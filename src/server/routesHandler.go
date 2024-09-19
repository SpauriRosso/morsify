package server

import "net/http"

func VerifyMethod(r *http.Request, method string) bool {
	return r.Method == method
}

func (s *Server) MosrsifyHandler(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, http.MethodPost) {
		http.Error(w, "Error: Invalid method", http.StatusMethodNotAllowed)
		return
	}

}

func (s *Server) DemorsifyHandler(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, http.MethodPost) {
		http.Error(w, "Error: Invalid method", http.StatusMethodNotAllowed)
		return
	}
}
