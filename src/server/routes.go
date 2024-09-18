package server

func (s *Server) Routes() {
	s.mux.HandleFunc("/morsify", s.MosrsifyHandler)
	s.mux.HandleFunc("/demorsify", s.DemorsifyHandler)
}
