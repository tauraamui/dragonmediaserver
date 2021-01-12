package web

import (
	"html/template"
	"net/http"
	"sync"
)

type TemplateData struct {
	Title   string
	Content template.HTML
}

func (s *Server) handleLoginGet() http.HandlerFunc {

	var (
		init         sync.Once
		templateHTML string
		pageHTML     string
		data         TemplateData
		tmpl         *template.Template
		err          error
	)

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			templateHTML, err = s.htmlRiceBox.String("template.html")
			if err != nil {
				s.errlog.Printf("Unable to load template HTML: %v\n", err)
			}

			pageHTML, err = s.htmlRiceBox.String("login/index.html")
			if err != nil {
				s.errlog.Printf("Unable to load login page HTML: %v\n", err)
			}

			data = TemplateData{
				Title:   "Login",
				Content: template.HTML(pageHTML),
			}

			tmpl, err = template.New("template").Parse(templateHTML)
			if err != nil {
				s.errlog.Printf("Unable to load template with login page content: %v\n", err)
			}
		})

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Missing HTML for login page."))
			return
		}

		if tmpl != nil {
			tmpl.Execute(w, data)
		}
	}
}
