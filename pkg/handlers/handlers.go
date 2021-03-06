package handlers

import (
	"net/http"

	"github.com/olumidesan/bookings/pkg/config"
	"github.com/olumidesan/bookings/pkg/models"
	"github.com/olumidesan/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: map[string]string{"Hello": "World", "remote_ip": remoteIP},
	})
}
