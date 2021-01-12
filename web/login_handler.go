package web

import (
	"net/http"
	"sync"
)

func (s *Server) handleLoginGet() http.HandlerFunc {

	var (
		init sync.Once
		p    string
		err  error
	)

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			p, err = s.htmlRiceBox.String("login/index.html")
			if err != nil {
				s.errlog.Printf("Unable to load login page HTML")
			}
		})

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Missing HTML for login page."))
			return
		}

		w.Write([]byte(p))
	}
}
