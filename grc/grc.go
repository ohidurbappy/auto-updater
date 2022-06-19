package grc 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Asset structure from GitHub API
type Asset struct {
	ID                 int       `json:"id"`
	URL                string    `json:"url"`
	Name               string    `json:"name"`
	Size               int       `json:"size"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}

// Release structure from GitHub API
type Release struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	AssetURL    string    `json:"asset_url"`
	UploadURL   string    `json:"upload_url"`
	HTMLURL     string    `json:"html_url"`
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	Prerelease  bool      `json:"prerelease"`
	Assets      []Asset   `json:"assets"`
}

const githubReleaseURLTemplate string = "https://api.github.com/repos/%s/%s/releases/latest"

// Check GitHub's Release API endpoint for the latest release (tag name)
// and returns with true or false as has it an update or don't
// and the latest *release as a struct
// and an error if something went wrong
func Check(repoOwner string, repoName string, currentVersion string) (bool, *Release, error) {
	fmt.Println("Check...")
	resp, err := http.Get(fmt.Sprintf(githubReleaseURLTemplate, repoOwner, repoName))
	if err != nil {
		return false, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err
	}

	releaseInfo := &Release{}
	err = json.Unmarshal(body, releaseInfo)
	if err != nil {
		return false, nil, err
	}

	if releaseInfo.TagName != currentVersion {
		return true, releaseInfo, nil
	}

	return false, releaseInfo, nil
}