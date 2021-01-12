package web

func (s *Server) routes() {
	s.router.ServeFiles("/public/*filepath", s.publicRiceBox.HTTPBox())
	s.router.HandlerFunc("GET", "/login", s.handleLoginGet())
}
