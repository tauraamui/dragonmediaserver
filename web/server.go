package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Server struct {
	dbConn *gorm.DB
	router *httprouter.Router
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{
		dbConn: db,
		router: httprouter.New(),
	}
	server.routes()
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
