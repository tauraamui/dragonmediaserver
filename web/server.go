package web

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/julienschmidt/httprouter"
	"github.com/tauraamui/dragonmediaserver/config"
	"gorm.io/gorm"
)

type Server struct {
	stdlog, errlog *log.Logger
	dbConn         *gorm.DB
	htmlRiceBox    *rice.Box
	publicRiceBox  *rice.Box
	router         *httprouter.Router
	cameras        []config.Camera
}

func NewServer(
	stdlog, errlog *log.Logger,
	db *gorm.DB,
	htmlRiceBox, publicRiceBox *rice.Box,
) *Server {
	cfg := config.LoadDragonDaemonConfig(stdlog, errlog)
	server := Server{
		stdlog:        stdlog,
		errlog:        errlog,
		dbConn:        db,
		htmlRiceBox:   htmlRiceBox,
		publicRiceBox: publicRiceBox,
		router:        httprouter.New(),
		cameras:       cfg.Cameras,
	}
	server.routes()
	return &server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

type response struct {
	body string
	code int
}

func (r *response) write(w http.ResponseWriter) {
	w.WriteHeader(r.code)
	w.Write([]byte(r.body))
}
