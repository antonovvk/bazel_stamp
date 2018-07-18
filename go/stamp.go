package stamp

import (
	"fmt"
)

var (
	StableGitTag    = "UNDEFINED"
	StableBuildDate = "UNDEFINED"
	BuildUser       = "UNDEFINED"
	BuildHost       = "UNDEFINED"
)

func Version() string {
	return fmt.Sprintf("Version: %s %s %s@%s", StableGitTag, StableBuildDate, BuildUser, BuildHost)
}
