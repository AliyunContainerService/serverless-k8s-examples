package version

import (
	"fmt"
)

type Info struct {
	GitCommit  string `json:"gitCommit" yaml:"gitCommit"`
	GitVersion string `json:"gitVersion" yaml:"gitVersion"`

	Message string `json:"message" yaml:"message"`
}

func (info Info) Info(name string) string {
	commit := info.GitCommit
	if len(commit) >= 8 {
		commit = info.GitCommit[:8]
	}
	return fmt.Sprintf("%s version %s, build %s/%s", name, info.GitVersion, commit, info.GitVersion)
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Version() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
	}
}
