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
	// +default=false
	dryRun bool,
	// +optional
	// +default=false
	ci bool,
	// +optional
	// +default=false
	debugMode bool,
) (string, error) {
	opts := dagger.SemanticReleaseRunOpts{
		DryRun:    dryRun,
		Ci:        ci,
		DebugMode: debugMode,
	}

	if len(repositoryUrl) > 0 {
		opts.RepositoryURL = repositoryUrl
	}

	secretEnvVarName := "GITHUB_TOKEN"

	return dag.SemanticRelease().Run(ctx, m.Worktree, secretEnvVarName, githubToken, opts)
}
