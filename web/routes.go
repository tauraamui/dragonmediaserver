package web

func (s *Server) routes() {
	s.router.ServeFiles("/public/*filepath", s.publicRiceBox.HTTPBox())
	s.router.HandlerFunc("GET", "/login", s.handleLoginGet())
	s.router.HandlerFunc("GET", "/cameras", s.handleCamerasGet())
	s.router.HandlerFunc("GET", "/camera/:title/manifest", s.handleCameraClipManifestGet())
}
