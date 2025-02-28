package version

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	Version   = "unknown-version"
	GitCommit = "unknown-commit"
	BuildDate = "unknown-builddate"
)

var versionTemplate = `Version:     {{.Version}}
Git commit:  {{.GitCommit}}{{if eq .GitState "dirty"}}
Git State:   {{.GitState}}{{end}}
Built:       {{.BuildDate}}
Go version:  {{.GoVersion}}
OS/Arch:     {{.OS}}/{{.Arch}}
`

type VersionInfo struct {
	Version   string
	GoVersion string
	GitCommit string
	GitState  string
	BuildDate string
	OS        string
	Arch      string
}

func New() *VersionInfo {
	i, err := strconv.ParseInt(BuildDate, 10, 64)
	if err != nil {
		panic(err)
	}

	tu := time.Unix(i, 0)

	vi := VersionInfo{
		Version:   Version,
		GoVersion: runtime.Version(),
		GitCommit: GitCommit,
		BuildDate: tu.String(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}

	if strings.HasSuffix(Version, "-dirty") {
		vi.GitState = "dirty"
	}

	return &vi
}

func (i *VersionInfo) Show() {
	tmpl, err := template.New("version").Parse(versionTemplate)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, i); err != nil {
		panic(err)
	}
}
