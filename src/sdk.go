package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"runtime"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/mholt/archiver"
)

const (
	golangDownloadURL  string = "https://dl.google.com/go"
	sdkDirName         string = "sdk"
	sdkUsedDirName     string = "current"
	sdkDownloadDirName string = ".downloads"
)

type (
	sdk struct {
		path     string
		download string
	}

	// ISDK is interface of SDK function
	ISDK interface {
		HasValidVersion(version string) bool
		Install(version string) error
		Switch(version string) error
		Show(version string)
		ShowList() error
		Uninstall(version string, purge bool) error
		Info() error
	}

	// SDK is information of Golang SDK
	SDK struct {
		Name           string
		Downloaded     bool
		Installed      bool
		Used           bool
		DownloadedPath string
		InstalledPath  string
	}
)

// SDK is functionality library
func (gvm gvm) SDK() ISDK {
	return sdk{
		path:     path.Join(gvm.path, sdkDirName),
		download: path.Join(gvm.path, sdkDownloadDirName),
	}
}

func (sdk sdk) HasValidVersion(version string) bool {
	return regexp.MustCompile(`^(\d+\.)(\d+\.)?(\*|\d+)$`).MatchString(version)
}

func (sdk sdk) List() ([]*SDK, error) {
	var sdks []*SDK
	files, err := ioutil.ReadDir(path.Join(sdk.path))
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			sdks = append(sdks, sdk.Get(f.Name()))
		}
	}
	return sdks, nil
}

func (sdk sdk) ShowList() error {
	sdks, err := sdk.List()
	if err != nil {
		return err
	}
	for _, s := range sdks {
		sdk.Show(s.Name)
		fmt.Println()
	}
	return nil
}

func (sdk sdk) Get(version string) *SDK {
	return &SDK{
		Name:           version,
		Downloaded:     sdk.HasDownloaded(version),
		Installed:      sdk.HasInstalled(version),
		Used:           sdk.HasUsed(version),
		DownloadedPath: sdk.DownloadedPath(version),
		InstalledPath:  sdk.InstalledPath(version),
	}
}

func (sdk sdk) Show(version string) {
	s := sdk.Get(version)
	fmt.Printf("Name:\t\t%s\n", s.Name)
	fmt.Printf("Used:\t\t%t\n", s.Used)
	fmt.Printf("Downloaded:\t%t", s.Downloaded)
	if s.Downloaded {
		fmt.Printf(" (%s)", s.DownloadedPath)
	}
	fmt.Println()
	fmt.Printf("Installed:\t%t", s.Installed)
	if s.Installed {
		fmt.Printf(" (%s)", s.InstalledPath)
	}
	fmt.Println()
}

func (sdk sdk) ArchiveName(version string) (string, error) {
	switch runtime.GOOS {
	case "darwin", "linux", "freebsd":
		return fmt.Sprintf("go%s.%s-amd64.tar.gz", version, runtime.GOOS), nil
	case "windows":
		return fmt.Sprintf("go%s.%s-amd64.zip", version, runtime.GOOS), nil
	default:
		return "", fmt.Errorf("your system is not compatible")
	}
}

func (sdk sdk) HasDownloaded(version string) bool {
	name, err := sdk.ArchiveName(version)
	if err != nil {
		return false
	}
	if _, err := os.Stat(path.Join(sdk.download, name)); os.IsNotExist(err) {
		return false
	}
	return true
}

func (sdk sdk) DownloadedPath(version string) string {
	name, err := sdk.ArchiveName(version)
	if err == nil {
		if sdk.HasDownloaded(version) {
			return path.Join(sdk.download, name)
		}
	}
	return ""
}

func (sdk sdk) Download(version string) (string, error) {
	name, err := sdk.ArchiveName(version)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s", golangDownloadURL, name)
	destination := path.Join(sdk.download, name)
	if !sdk.HasDownloaded(version) {
		client := grab.NewClient()
		req, _ := grab.NewRequest(fmt.Sprintf("%s.tmp", destination), url)
		fmt.Printf("Downloading %v...", req.URL())
		resp := client.Do(req)
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()
	Loop:
		for {
			select {
			case <-t.C:
				fmt.Printf("\rDownloading %v... %.2f%%", req.URL(), 100*resp.Progress())
			case <-resp.Done:
				break Loop
			}
		}
		if err := resp.Err(); err != nil {
			fmt.Printf("\rDownloading %v... failed\n", req.URL())
			return "", err
		}
		if resp.IsComplete() {
			if err := os.Rename(fmt.Sprintf("%s.tmp", destination), destination); err != nil {
				fmt.Printf("\rDownloading %v... failed\n", req.URL())
				return "", err
			}
		}
		fmt.Printf("\rDownloading %v... success\n", req.URL())
	} else {
		fmt.Printf("\rDownloading %v... use cache\n", url)
	}
	return destination, nil
}

func (sdk sdk) HasInstalled(version string) bool {
	if _, err := os.Stat(path.Join(sdk.path, version)); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (sdk sdk) InstalledPath(version string) string {
	if sdk.HasInstalled(version) {
		return path.Join(sdk.path, version)
	}
	return ""
}

func (sdk sdk) Install(version string) error {
	if !sdk.HasInstalled(version) {
		download, err := sdk.Download(version)
		if err != nil {
			return err
		}
		fmt.Printf("Installing SDK %s...", version)
		dir, err := ioutil.TempDir("", version)
		if err != nil {
			fmt.Printf("\rInstalling SDK %s... failed", version)
			return err
		}
		defer os.RemoveAll(dir)
		if err := archiver.Unarchive(download, dir); err != nil {
			fmt.Printf("\rInstalling SDK %s... failed", version)
			return err
		}
		if err := os.MkdirAll(sdk.path, 0755); err != nil {
			fmt.Printf("\rInstalling SDK %s... failed", version)
			return err
		}
		if err := os.Rename(path.Join(dir, "go"), path.Join(sdk.path, version)); err != nil {
			fmt.Printf("\rInstalling SDK %s... failed", version)
			return err
		}
		fmt.Printf("\rInstalling SDK %s... success\n", version)
		return nil
	}
	fmt.Printf("SDK %s already installed\n", version)
	return nil
}

func (sdk sdk) Unset() error {
	if _, err := os.Lstat(path.Join(sdk.path, sdkUsedDirName)); err == nil {
		if err := os.Remove(path.Join(sdk.path, sdkUsedDirName)); err != nil {
			return fmt.Errorf("failed to unlink: %+v", err)
		}
	}
	return nil
}

func (sdk sdk) Switch(version string) error {
	if !sdk.HasInstalled(version) {
		return fmt.Errorf("version not installed")
	}
	if err := sdk.Unset(); err != nil {
		return err
	}
	if err := os.Symlink(path.Join(sdk.path, version), path.Join(sdk.path, sdkUsedDirName)); err != nil {
		return fmt.Errorf("failed to create symlink: %+v", err)
	}
	fmt.Printf("Switched to SDK %s\n", version)
	return nil
}

func (sdk sdk) HasUsed(version string) bool {
	usedVersion, _ := sdk.Used()
	return usedVersion == version
}

func (sdk sdk) Used() (string, error) {
	if _, err := os.Stat(path.Join(sdk.path, sdkUsedDirName)); !os.IsNotExist(err) {
		current, err := os.Readlink(path.Join(sdk.path, sdkUsedDirName))
		if err != nil {
			return "", err
		}
		return path.Base(current), nil
	}
	return "", fmt.Errorf("no version used")
}

func (sdk sdk) Info() error {
	used, err := sdk.Used()
	if err != nil {
		return err
	}
	sdk.Show(used)
	return nil
}

func (sdk sdk) Uninstall(version string, purge bool) error {
	if sdk.HasUsed(version) {
		fmt.Printf("Unlink used SDK...")
		if err := sdk.Unset(); err != nil {
			fmt.Printf(" failed\n")
			return err
		}
		fmt.Printf(" success\n")
	}
	if sdk.HasInstalled(version) {
		fmt.Printf("Uninstall SDK %s ...", version)
		if err := os.RemoveAll(sdk.InstalledPath(version)); err != nil {
			fmt.Printf(" failed\n")
			return err
		}
		fmt.Printf(" success\n")
	}
	if purge {
		fmt.Printf("Remove SDK %s ...", version)
		if sdk.HasDownloaded(version) {
			if err := sdk.Remove(version); err != nil {
				fmt.Printf(" failed\n")
				return err
			}
		}
		fmt.Printf(" success\n")
	}
	return nil
}

func (sdk sdk) Remove(version string) error {
	return os.RemoveAll(sdk.DownloadedPath(version))
}
