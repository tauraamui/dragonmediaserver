package web

func (s *Server) routes() {
	s.router.GET("/login", s.handleLoginGet())
}
