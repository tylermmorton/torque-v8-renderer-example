package login

import (
	"io/fs"
	"net/http"

	"github.com/tylermmorton/torque"
)

type ViewModel struct {
	Message string `json:"message"`
}

type Controller struct {
	Dist fs.FS
}

var _ interface {
	torque.Loader[ViewModel]
} = &Controller{}

func (m *Controller) Plugins() []torque.Plugin {
	clientBuild, err := fs.Sub(m.Dist, ".dist/client")
	if err != nil {
		panic(err)
	}

	serverBuild, err := fs.Sub(m.Dist, ".dist/server")
	if err != nil {
		panic(err)
	}

	return []torque.Plugin{
		&torque.VitePlugin{
			ClientBuild: clientBuild,
			ServerBuild: serverBuild,
		},
	}
}

func (m *Controller) Load(req *http.Request) (ViewModel, error) {
	return ViewModel{
		Message: "Hello, World!",
	}, nil
}
