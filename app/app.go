package app

import (
	"embed"
	_ "embed"
	"net/http"

	"github.com/tylermmorton/torque"
	"github.com/tylermmorton/torque-v8-renderer-example/app/routes/login"
)

//go:embed .dist
var dist embed.FS

//go:embed app.tmpl.html
var templateText string

type ViewModel struct {
	Title string
}

func (ViewModel) TemplateText() string {
	return templateText
}

type Controller struct{}

var _ interface {
	torque.Loader[ViewModel]
	torque.RouterProvider
} = &Controller{}

func (m *Controller) Router(r torque.Router) {
	r.Handle("/s", http.FileServer(http.FS(dist)))
	r.Handle("/login", torque.MustNew[login.ViewModel](&login.Controller{Dist: dist}))
}

func (m *Controller) Load(req *http.Request) (ViewModel, error) {
	return ViewModel{
		Title: torque.UseTitle(req),
	}, nil
}
