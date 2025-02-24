package main

import (
	"context"
	"fmt"
	"log"

	"dagger/mybar-barista/internal/dagger"
)

// Publish the built binaries as a GitHub release
func (m *MybarBuild) Publish(
	ctx context.Context,
	token *dagger.Secret,
) error {
	if token == nil {
		return fmt.Errorf("GitHub token is required for publishing")
	}

	if m.Git.Tag == "" {
		return fmt.Errorf("Git tag is missing, cannot create a release")
	}

	filesStr, err := m.BinDir.Entries(ctx)
	if err != nil {
		return fmt.Errorf("failed to list build directory entries: %w", err)
	}

	var files []*dagger.File
	for _, f := range filesStr {
		files = append(files, m.BinDir.File(f))
	}

	if len(files) == 0 {
		return fmt.Errorf("no files found to publish")
	}

	opts := dagger.GhReleaseCreateOpts{
		Files:     files,
		VerifyTag: true,
		Token:     token,
		Repo:      appSourceUrl,
	}

	log.Printf("Publishing release for tag %s with %d files...\n", m.Git.Tag, len(files))
	err = dag.Gh().Release().Create(ctx, m.Git.Tag, appName, opts)
	if err != nil {
		return fmt.Errorf("failed to create GitHub release: %w", err)
	}

	log.Println("Release published successfully!")
	return nil
}
