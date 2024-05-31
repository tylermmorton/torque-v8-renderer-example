package login

import (
	"io/fs"
	"log"
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

func logFileSystem(fsys fs.FS) {
	var walkFn func(path string, d fs.DirEntry, err error) error

	walkFn = func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if d.IsDir() {
			log.Printf("Dir: %s", path)
		} else {
			log.Printf("File: %s", path)
		}
		return nil
	}

	err := fs.WalkDir(fsys, ".", walkFn)
	if err != nil {
		panic(err)
	}
}

func (m *Controller) Plugins() []torque.Plugin {
	logFileSystem(m.Dist)

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
	return ViewModel{}, nil
}
