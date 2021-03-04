package emvers

import _ "embed"

//go:embed versionfile
var version string

func GitVersion() string {
	return version
}
