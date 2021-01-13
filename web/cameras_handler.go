package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleCamerasGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleCameraClipManifestGet() http.HandlerFunc {
	manifestCache := map[string]string{}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			body string
			code int
		}{
			"",
			404,
		}

		defer func() {
			w.WriteHeader(resp.code)
			w.Write([]byte(resp.body))
		}()

		urlParams := httprouter.ParamsFromContext(r.Context())
		cameraTitle := urlParams.ByName("title")
		if manifest, ok := manifestCache[cameraTitle]; ok {
			resp.body = manifest
			resp.code = 200
			return
		}

		for _, camera := range s.cameras {
			if camera.Title == cameraTitle {
				manifest := "MANIFEST"
				manifestCache[cameraTitle] = manifest
				resp.body = manifest
				resp.code = 200
				return
			}
		}
	}
}
