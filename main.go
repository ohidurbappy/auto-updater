package main

  import (
	"fmt"
  )

import GRC "github.com/ohidurbappy/auto-updater/github-release-checker"

const (
	RepoOwner string = "ohidurbappy"
	RepoName  string = "auto-updater"
	Version   string = "v1.0"
  )

  func main() {
	needUpdate, releaseInfo, err := GRC.Check(RepoOwner, RepoName, Version)
  
	if err != nil {
	  panic(err)
	}
  
	if needUpdate {
	  fmt.Printf("Please update from %s to %s at %s", Version, releaseInfo.TagName, releaseInfo.HTMLURL)
	} else {
	  fmt.Println("Application status: Up to date.")
	}
  }