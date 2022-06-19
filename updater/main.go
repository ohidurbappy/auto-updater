package updater

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"path"
	_ "runtime"

	"github.com/kardianos/osext"
	grc "github.com/ohidurbappy/auto-updater/grc"
	AppInfo "github.com/ohidurbappy/auto-updater/info"
)


// Update structure is the representation of the update command.
type Update struct{}

const (
	binaryChmodValue = 0o755
)

// opts *commander.CommandHelper
// Execute is the main function. It will be called on update command.
func (c *Update) Execute() {
	hasUpdate, release, _ := grc.Check(AppInfo.RepoOwner, AppInfo.Name, AppInfo.Version)

	if !hasUpdate {
		fmt.Printf("You are using the latest version of %s \nVersion: %s\n", AppInfo.Name, AppInfo.Version)

		return
	}

	var (
		assetToDownload grc.Asset
		found           bool
	)

	for _, asset := range release.Assets {
		if asset.Name == c.buildFilename(release.TagName) {
			assetToDownload = asset
			found = true

			break
		}
	}

	if !found {
		fmt.Printf("Your %s is up-to-date. \\o/\n", AppInfo.Name)

		return
	}

	downloadError := c.downloadBinary(assetToDownload.BrowserDownloadURL)
	if downloadError != nil {
		fmt.Printf("Error: %s\n", downloadError.Error())
	}

	fmt.Printf("Now you have a fresh new %s\n", AppInfo.Name)
}

func (c *Update) buildFilename(version string) string {
	// return fmt.Sprintf("%s-%s-%s-%s.tar.gz", AppInfo.Name, version, runtime.GOOS, runtime.GOARCH)
	return fmt.Sprintf("%s.tar.gz", AppInfo.Name)
}

func (c *Update) downloadBinary(uri string) error {
	fmt.Println(" -> Download...")

	client := http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), "GET", uri, nil)
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	response, err := client.Do(request)
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	defer response.Body.Close()

	gzipReader, _ := gzip.NewReader(response.Body)
	defer gzipReader.Close()

	fmt.Println(" -> Extract...")
	tarReader := tar.NewReader(gzipReader)

	_, err = tarReader.Next()
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	currentExecutable, _ := osext.Executable()
	originalPath := path.Dir(currentExecutable)

	file, err := ioutil.TempFile(originalPath, AppInfo.Name)
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	defer file.Close()

	_, err = io.Copy(file, tarReader) //nolint:gosec // I don't have better option right now.
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	err = file.Chmod(binaryChmodValue)
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	// first rename the running application
	err = os.Rename(currentExecutable, path.Join(originalPath, AppInfo.Name+"_old"+".exe"))
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	// sleep for a while to let the old application finish its work
	time.Sleep(time.Second * 2)

	// then move the new binary to the original path
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	err = ioutil.WriteFile(currentExecutable, content, binaryChmodValue)
	if err != nil {
		return DownloadError{Message: err.Error()}
	}

	// finally remove the temp file
	err = os.Remove(file.Name())
	if err != nil {
		return DownloadError{Message: err.Error()}
	}



	return nil
}
