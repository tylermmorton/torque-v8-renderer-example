package app

import (
	"embed"
	_ "embed"
	"io/fs"
	"net/http"

	"github.com/tylermmorton/torque"
	"github.com/tylermmorton/torque-v8-renderer-example/app/routes/login"
)

//go:embed .dist
var dist embed.FS

//go:embed app.tmpl.html
var templateText string

type ViewModel struct {
	Title   string
	Scripts []string
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
	clientDist, err := fs.Sub(dist, ".dist/client")
	if err != nil {
		panic(err)
	}

	r.HandleFileSystem("/s", clientDist)
	r.Handle("/login", torque.MustNew[login.ViewModel](&login.Controller{Dist: dist}))
}

func (m *Controller) Load(req *http.Request) (ViewModel, error) {
	return ViewModel{
		Title:   torque.UseTitle(req),
		Scripts: torque.UseScripts(req),
	}, nil
}
