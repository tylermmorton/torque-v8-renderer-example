package login

import (
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/tylermmorton/torque"
	"github.com/tylermmorton/torque/pkg/plugins/v8"
	"github.com/tylermmorton/torque/pkg/vite"
)

type ViewModel struct {
	Count   int    `json:"count"`
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

	clientManifest, err := vite.ParseManifestFromFS(clientBuild, "manifest.json")
	if err != nil {
		panic(err)
	}

	clientDist := v8.Dist{
		Dist: clientBuild,
		ResolverFn: func(s string) (string, error) {
			if manifestEntry, ok := clientManifest[s]; ok {
				return filepath.Join("/s/", manifestEntry.File), nil
			}
			return s, fmt.Errorf("entry %s not found in vite manifest", s)
		},
	}

	serverBuild, err := fs.Sub(m.Dist, ".dist/server")
	if err != nil {
		panic(err)
	}

	serverManifest, err := vite.ParseManifestFromFS(serverBuild, "manifest.json")
	if err != nil {
		panic(err)
	}

	serverDist := v8.Dist{
		Dist: serverBuild,
		ResolverFn: func(s string) (string, error) {
			if manifestEntry, ok := serverManifest[s]; ok {
				return manifestEntry.File, nil
			}
			return s, fmt.Errorf("entry %s not found in vite manifest", s)
		},
	}

	return []torque.Plugin{
		v8.NewPlugin(&serverDist, &clientDist),
	}
}

func (m *Controller) Load(req *http.Request) (ViewModel, error) {
	return ViewModel{
		Count:   5,
		Message: "Hello, World!",
	}, nil
}
