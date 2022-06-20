package main

  import (
	"fmt"
	"os"
    "os/exec"
    "runtime"
    "sync"
	"syscall"
	"time"
    
	GRC "github.com/ohidurbappy/auto-updater/grc"
	AppInfo "github.com/ohidurbappy/auto-updater/info"
	Updater "github.com/ohidurbappy/auto-updater/updater"
  )

  var wg sync.WaitGroup

  func main() {
	fmt.Printf("Current Version: %s\n", AppInfo.Version)
	
	
	hasUpdate, releaseInfo, err := GRC.Check(AppInfo.RepoOwner, AppInfo.Name, AppInfo.Version)
  
	if err != nil {
	  panic(err)
	}
  
	if hasUpdate {
	 
		fmt.Printf("Please update from %s to %s at %s", AppInfo.Version, releaseInfo.TagName, releaseInfo.HTMLURL)
		var updater=&Updater.Update{}
		updater.Execute()
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println("Restarting...")
			RestartSelf()
		}()
		wg.Wait()
		
	  } else {
		fmt.Printf("You are using the latest version of %s \nVersion: %s\n", AppInfo.Name, AppInfo.Version)
	}

	for {
		fmt.Println("Doing the real work here...")
		time.Sleep(time.Second * 5)
	}

	


  }



  func RestartSelf() error {
    self, err := os.Executable()
    if err != nil {
        return err
    }
    args := os.Args
    env := os.Environ()
    // Windows does not support exec syscall.
    if runtime.GOOS == "windows" {
        cmd := exec.Command(self, args[1:]...)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Stdin = os.Stdin
        cmd.Env = env
        err := cmd.Run()
        if err == nil {
            os.Exit(0)
        }
        return err
    }
    return syscall.Exec(self, args, env)
}
