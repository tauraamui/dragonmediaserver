package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
	"github.com/tauraamui/dragonmediaserver/config"
)

func (s *Server) handleCamerasGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleCameraClipManifestGet() http.HandlerFunc {
	manifestCache := map[string]string{}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := response{}
		defer resp.write(w)

		urlParams := httprouter.ParamsFromContext(r.Context())
		cameraTitle := urlParams.ByName("title")
		date := urlParams.ByName("date")
		manifestKey := fmt.Sprintf("%s|%s", cameraTitle, date)

		if manifest, ok := manifestCache[manifestKey]; ok {
			s.stdlog.Println("Returning cached manifest version")
			resp.body = manifest
			resp.code = 200
			return
		}

		for _, camera := range s.cameras {
			if camera.Title == cameraTitle {
				manifest, err := generateManifest(camera, date)
				if err != nil {
					resp.body = fmt.Sprintf("Unable to generate manifest: %s", err)
					resp.code = 500
					return
				}
				manifestCache[manifestKey] = manifest
				resp.body = manifest
				resp.code = 200
				return
			}
		}

	}
}

func generateManifest(camera config.Camera, date string) (string, error) {
	rootPath := fmt.Sprintf("%s/%s/%s", camera.PersistLoc, camera.Title, date)
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return "", err
	}

	cmdArgs := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		escapedFileName := f.Name()
		cmdArgs = append(cmdArgs, fmt.Sprintf("in=%s/%s,stream=video,output=%s", rootPath, escapedFileName, escapedFileName))
	}

	if len(files) > 0 {
		cmdArgs = append(cmdArgs, "--mpd_output h264.mpd")
	}

	fmt.Printf("ARGS LIST: %s", cmdArgs)
	cmd := exec.Command("./packager-osx", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	return "", err
}
