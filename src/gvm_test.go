package src

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	gvmTest = gvm{
		path: path.Join(os.Getenv("GOPATH"), "src/github.com/tfournier/gvm/.test"),
	}
)

func TestGvm_Config(t *testing.T) {
	cfg := gvmTest.Config()
	assert.Equal(t, path.Join(gvmTest.path, "sdk", "current"), cfg["GOROOT"])
}
