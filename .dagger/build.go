package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"dagger/mybar-barista/internal/dagger"

	cplatforms "github.com/containerd/platforms"
	"github.com/juli3nk/go-utils/ci"
)

type MybarBuild struct {
	Git      *Git
	BinFiles []*dagger.File
}

// Build binaries
func (m *MybarBarista) Build(
	ctx context.Context,
	// +optional
	version string,
) (*MybarBuild, error) {
	platformSpecifiers := []string{
		"linux/amd64",
	}
	platforms, err := cplatforms.ParseAll(platformSpecifiers)
	if err != nil {
		return nil, err
	}

	appVersion := ci.ResolveVersion(version, m.Git.Tag, m.Git.Commit, m.Git.Uncommitted)
	goAppVersionPkgPath := fmt.Sprintf("%s/pkg/version", appSourceUrl)
	tsNow := time.Now()

	goBuildPackages := []string{"."}
	goBuildLdflags := []string{
		fmt.Sprintf("-X %s.Version=%s", goAppVersionPkgPath, appVersion),
		fmt.Sprintf("-X %s.GitCommit=%s", goAppVersionPkgPath, m.Git.Commit),
		fmt.Sprintf("-X %s.BuildDate=%d", goAppVersionPkgPath, tsNow.Unix()),
	}

	var wg sync.WaitGroup
	errorsChan := make(chan error, len(platforms))

	output := []*dagger.File{}

	for _, platform := range platforms {
		wg.Add(1)
		go func(platform cplatforms.Platform) {
			defer wg.Done()
			opts := dagger.GoBuildOpts{
				CgoEnabled: "1",
				Ldflags:    goBuildLdflags,
				Musl:       true,
				Arch:       platform.Architecture,
				Os:         platform.OS,
			}
			goBuilder := dag.Go(goVersion, m.Worktree).Build(appName, goBuildPackages, opts)

			output = append(output, goBuilder)
			errorsChan <- nil
		}(platform)
	}

	wg.Wait()
	close(errorsChan)

	var buildErrors []error
	for err := range errorsChan {
		if err != nil {
			buildErrors = append(buildErrors, err)
		}
	}

	if len(buildErrors) > 0 {
		return nil, fmt.Errorf("build failed: %w", errors.Join(buildErrors...))
	}

	return &MybarBuild{
		Git:      m.Git,
		BinFiles: output,
	}, nil
}

func (m *MybarBuild) Stdout(ctx context.Context) (string, error) {
	outputs := []string{}

	for _, file := range m.BinFiles {
		name, err := file.Name(ctx)
		if err != nil {
			return "", err
		}

		outputs = append(outputs, name)
	}

	return strings.Join(outputs, "\n"), nil
}

func (m *MybarBuild) Dir(ctx context.Context) (*dagger.Directory, error) {
	dir := dag.Directory()

	for _, file := range m.BinFiles {
		name, err := file.Name(ctx)
		if err != nil {
			return nil, err
		}

		dir = dir.WithFile(name, file)
	}

	return dir, nil
}
