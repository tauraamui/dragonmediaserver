package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tauraamui/dragonmediaserver/config"
)

func (s *Server) handleCameraClipManifestGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlParams := httprouter.ParamsFromContext(r.Context())

		cameraTitle := urlParams.ByName("title")

		var matchedCamera *config.Camera
		for _, camera := range s.cameras {
			if camera.Title == cameraTitle {
				matchedCamera = &camera
			}
		}

		if matchedCamera == nil {
			w.WriteHeader(404)
			return
		}

		w.Write([]byte(matchedCamera.PersistLoc))
	}
}
