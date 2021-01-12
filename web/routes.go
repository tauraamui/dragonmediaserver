package web

func (s *Server) routes() {
	s.router.HandlerFunc("GET", "/login", s.handleLoginGet())
}
