package app

import (
	_ "embed"
	"io/fs"
	"net/http"

	"github.com/tylermmorton/torque"
	"github.com/tylermmorton/torque-v8-renderer-example/app/routes/login"
)

//go:embed app.tmpl.html
var templateText string

type ViewModel struct {
	Title string
}

func (ViewModel) TemplateText() string {
	return templateText
}

type Controller struct {
	Dist fs.FS
}

var _ interface {
	torque.Loader[ViewModel]
	torque.RouterProvider
} = &Controller{}

func (m *Controller) Router(r torque.Router) {
	r.Handle("/s", http.FileServer(http.FS(m.Dist)))
	r.Handle("/login", torque.MustNew[login.ViewModel](&login.Controller{Dist: m.Dist}))
}

func (m *Controller) Load(req *http.Request) (ViewModel, error) {
	return ViewModel{
		Title: torque.UseTitle(req),
	}, nil
}
