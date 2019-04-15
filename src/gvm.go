package src

import (
	"fmt"
	"log"
	"os/user"
	"path"
)

const (
	gvmDirName string = ".gvm"
)

type (
	gvm struct {
		path string
	}

	// IGVM is interface of GVM function
	IGVM interface {
		SDK() ISDK
		ShowConfig()
	}
)

// GVM is functionality library
func GVM() IGVM {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return gvm{
		path: fmt.Sprintf("%s/%s", usr.HomeDir, gvmDirName),
	}
}

func (gvm gvm) ShowConfig() {
	for k, v := range gvm.Config() {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func (gvm gvm) Config() map[string]string {
	var e = make(map[string]string)
	e["GOROOT"] = path.Join(gvm.path, sdkDirName, sdkUsedDirName)
	e["PATH"] = "$PATH:" + path.Join(gvm.path, sdkDirName, sdkUsedDirName, "bin")
	return e
}
