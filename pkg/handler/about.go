package handler

import (
	"log"
	"net/http"

	"github.com/captv89/web-udemy-practice/pkg/models"
	"github.com/captv89/web-udemy-practice/pkg/render"
)

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("About Page")

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["title"] = "About"
	stringMap["body"] = "This is the about page :)"
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
