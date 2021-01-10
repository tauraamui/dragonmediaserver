package web

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Server struct {
	dbConn  *gorm.DB
	riceBox *rice.Box
	router  *httprouter.Router
}

func NewServer(db *gorm.DB, riceBox *rice.Box) *Server {
	server := &Server{
		dbConn:  db,
		riceBox: riceBox,
		router:  httprouter.New(),
	}
	server.routes()
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
