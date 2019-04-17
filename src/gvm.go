package src

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/spf13/cobra"
	"github.com/tfournier/completion"
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
		Initialize(cmd *cobra.Command) error
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

func (gvm gvm) Initialize(cmd *cobra.Command) error {
	if err := os.MkdirAll(gvm.path, 0755); err != nil {
		return err
	}
	if err := completion.Cobra(cmd).Zsh().ToFile(path.Join(gvm.path, "completion.zsh")); err != nil {
		return err
	}
	if err := completion.Cobra(cmd).Bash().ToFile(path.Join(gvm.path, "completion.bash")); err != nil {
		return err
	}
	fmt.Println("Configuration:")
	fmt.Printf("\tFor Zsh: `eval $(gvm config)` in ~/.zshrc\n")
	fmt.Printf("\tFor Bash: `eval $(gvm config)` in ~/.bashrc\n")
	fmt.Printf("\t\n")
	fmt.Println("Completion:")
	fmt.Printf("\tFor Zsh: `source %s` in ~/.zshrc\n", path.Join(gvm.path, "completion.zsh"))
	fmt.Printf("\tFor Bash: `source %s` in ~/.bashrc\n", path.Join(gvm.path, "completion.bash"))
	return nil
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
