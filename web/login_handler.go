package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleLoginGet() httprouter.Handle {
	loginHTML, _ := s.riceBox.String("login/index.html")
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte(loginHTML))
	}
}
