package info

import (
	_ "embed"
)

// RepoOwner defined the owner of the repo on GitHub.
const RepoOwner string = "ohidurbappy"

// Name defined the application name.
const Name string = "auto-updater"

// Version defined current version of this application.


//go:embed version.txt
var Version string