package delivery

import "net/http"

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Route() *http.ServeMux {
	s.mux.HandleFunc("/message", MessageHandler)
	return s.mux
}
