package utils

import (
	"fmt"
	"runtime"
)

var (
	version      string
	gitBranch    string
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
)

type Git struct {
	Branch    string `json:"branch,omitempty"`
	Tag       string `json:"tag,omitempty"`
	Commit    string `json:"commit,omitempty"`
	TreeState string `json:"tree_state,omitempty"`
}

type Info struct {
	Version   string `json:"version,omitempty"`
	Git       Git    `json:"git,omitempty"`
	BuildDate string `json:"buildAt,omitempty"`
	GoVersion string `json:"goVersion,omitempty"`
	Compiler  string `json:"compiler,omitempty"`
	Platform  string `json:"platform,omitempty"`
}

func (info Info) String() string {
	return info.Git.Commit
}

func Version() Info {
	return Info{
		Version: version,
		Git: Git{
			Branch:    gitBranch,
			Tag:       gitTag,
			Commit:    gitCommit,
			TreeState: gitTreeState,
		},
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
