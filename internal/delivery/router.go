package delivery

import (
	"gobot/internal/models"
	"net/http"
)

type Server struct {
	mux     *http.ServeMux
	channel chan models.Message
}

func New(ch chan models.Message) *Server {
	return &Server{
		mux:     http.NewServeMux(),
		channel: ch,
	}
}

func (s *Server) Route() *http.ServeMux {
	s.mux.HandleFunc("/message", s.MessageHandler)
	return s.mux
}
