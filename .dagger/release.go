package main

import (
	"context"

	"dagger/mybar-barista/internal/dagger"
)

// Release triggers a semantic release process for a repository
func (m *MybarBarista) Release(
	ctx context.Context,
	githubToken *dagger.Secret,
	// +optional
	repositoryUrl string,
	// +optional
	// +default=true
	dryRun bool,
	// +optional
	// +default=false
	ci bool,
	// +optional
	// +default=false
	debug bool,
) (string, error) {
	opts := dagger.SemanticReleaseRunOpts{}

	if len(repositoryUrl) > 0 {
		opts.RepositoryURL = repositoryUrl
	}
	if dryRun {
		opts.DryRun = true
	}
	if ci {
		opts.Ci = true
	}
	if debug {
		opts.Debug = true
	}

	secretEnvVarName := "GITHUB_TOKEN"

	return dag.SemanticRelease().Run(ctx, m.Worktree, secretEnvVarName, githubToken, opts)
}
