package web

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Server struct {
	stdlog, errlog *log.Logger
	dbConn         *gorm.DB
	htmlRiceBox    *rice.Box
	publicRiceBox  *rice.Box
	router         *httprouter.Router
}

func NewServer(
	stdlog, errlog *log.Logger,
	db *gorm.DB,
	htmlRiceBox, publicRiceBox *rice.Box,
) *Server {
	server := Server{
		stdlog:        stdlog,
		errlog:        errlog,
		dbConn:        db,
		htmlRiceBox:   htmlRiceBox,
		publicRiceBox: publicRiceBox,
		router:        httprouter.New(),
	}
	server.routes()
	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
