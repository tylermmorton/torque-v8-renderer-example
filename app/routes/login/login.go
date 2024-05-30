package login

import (
	"io/fs"
	"net/http"

	"github.com/tylermmorton/torque"
)

type ViewModel struct{}

type Controller struct {
	Dist fs.FS
}

var _ interface {
	torque.Loader[ViewModel]
} = &Controller{}

func (m *Controller) Plugins() ([]torque.Plugin, error) {
	clientBuild, err := fs.Sub(m.Dist, "client")
	if err != nil {
		return nil, err
	}

	serverBuild, err := fs.Sub(m.Dist, "server")
	if err != nil {
		return nil, err
	}

	return []torque.Plugin{
		&torque.VitePlugin{
			ClientBuild: clientBuild,
			ServerBuild: serverBuild,
		},
	}, nil
}

func (m *Controller) Load(req *http.Request) (ViewModel, error) {
	return ViewModel{}, nil
}
