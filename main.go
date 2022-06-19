package main

  import (
	"fmt"
	// GRC "github.com/ohidurbappy/auto-updater/grc"
	AppInfo "github.com/ohidurbappy/auto-updater/info"
	Updater "github.com/ohidurbappy/auto-updater/updater"
  )



  func main() {
	// hasUpdate, releaseInfo, err := GRC.Check(AppInfo.RepoOwner, AppInfo.Name, AppInfo.Version)
  
	// if err != nil {
	//   panic(err)
	// }
  
	// if hasUpdate {
	//   fmt.Printf("Please update from %s to %s at %s", AppInfo.Version, releaseInfo.TagName, releaseInfo.HTMLURL)
	
	
	//   } else {
	// 	fmt.Printf("You are using the latest version of %s \nVersion: %s\n", AppInfo.Name, AppInfo.Version)
	// }

	fmt.Printf("Current Version: %s\n", AppInfo.Version)
	var updater=&Updater.Update{}
	updater.Execute()

  }