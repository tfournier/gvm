package src

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sdkTest = sdk{
		path:     path.Join(gvmTest.path, "sdk"),
		download: path.Join(gvmTest.path, ".downloads"),
	}
	validSdkVersion   = "1.12.3"
	unValidSdkVersion = "1.12.0"
)

func TestGvm_SDK(t *testing.T) {
	assert.Equal(t, sdkTest, gvmTest.SDK())
}

func TestSdk_HasValidVersion(t *testing.T) {
	assert.Equal(t, true, sdkTest.HasValidVersion("1.12.3"))
	assert.Equal(t, true, sdkTest.HasValidVersion("1.12"))
	assert.Equal(t, false, sdkTest.HasValidVersion("1."))
	assert.Equal(t, false, sdkTest.HasValidVersion("1.12."))
	assert.Equal(t, false, sdkTest.HasValidVersion("test"))
}

func TestSdk_Install(t *testing.T) {

	// Not valid
	assert.Error(t, sdkTest.Install(unValidSdkVersion))

	unValidSdkName := fmt.Sprintf("go%s.%s-%s.tar.gz", unValidSdkVersion, runtime.GOOS, runtime.GOARCH)

	t.Log(path.Join(sdkTest.download, unValidSdkName))
	_, unValidSdkDownloaded := os.Stat(path.Join(sdkTest.download, unValidSdkName))
	assert.Equal(t, true, os.IsNotExist(unValidSdkDownloaded))

	t.Log(path.Join(sdkTest.path, unValidSdkName))
	_, unValidSdkInstalled := os.Stat(path.Join(sdkTest.path, unValidSdkName))
	assert.Equal(t, true, os.IsNotExist(unValidSdkInstalled))

	// Valid
	assert.NoError(t, sdkTest.Install(validSdkVersion))

	validSdkName := fmt.Sprintf("go%s.%s-%s.tar.gz", validSdkVersion, runtime.GOOS, runtime.GOARCH)

	t.Log(path.Join(sdkTest.download, validSdkName))
	_, validSdkDownloaded := os.Stat(path.Join(sdkTest.download, validSdkName))
	assert.Equal(t, false, os.IsNotExist(validSdkDownloaded))

	t.Log(path.Join(sdkTest.path, validSdkVersion))
	_, validSdkInstalled := os.Stat(path.Join(sdkTest.path, validSdkVersion))
	assert.Equal(t, false, os.IsNotExist(validSdkInstalled))
}

func TestSdk_Switch(t *testing.T) {
	// Not valid
	assert.Error(t, sdkTest.Switch(unValidSdkVersion))

	// Valid
	assert.NoError(t, sdkTest.Switch(validSdkVersion))
}

func TestSdk_HasDownloaded(t *testing.T) {
	assert.Equal(t, false, sdkTest.HasDownloaded(unValidSdkVersion))
	assert.Equal(t, true, sdkTest.HasDownloaded(validSdkVersion))
}

func TestSdk_HasInstalled(t *testing.T) {
	assert.Equal(t, false, sdkTest.HasInstalled(unValidSdkVersion))
	assert.Equal(t, true, sdkTest.HasInstalled(validSdkVersion))
}

func TestSdk_HasUsed(t *testing.T) {
	assert.Equal(t, false, sdkTest.HasUsed(unValidSdkVersion))
	assert.Equal(t, true, sdkTest.HasUsed(validSdkVersion))
}

func TestSdk_Get(t *testing.T) {
	s := sdkTest.Get(validSdkVersion)

	assert.Equal(t, validSdkVersion, s.Name)
	assert.Equal(t, true, s.Used)
	assert.Equal(t, true, s.Installed)
	assert.Equal(t, true, s.Downloaded)
}

func TestSdk_List(t *testing.T) {
	sdks, err := sdkTest.List()
	assert.NoError(t, err)
	fmt.Println(sdks)
	fmt.Println(len(sdks))
}

func TestSdk_Uninstall(t *testing.T) {

	// Valid
	assert.NoError(t, sdkTest.Uninstall(validSdkVersion, true))

}
