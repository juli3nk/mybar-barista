package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"dagger/mybar-barista/internal/dagger"
)

// Publish the built binaries as a GitHub release
func (m *MybarBuild) Publish(
	ctx context.Context,
	token *dagger.Secret,
) error {
	if len(m.BinFiles) == 0 {
		return fmt.Errorf("no files found to publish")
	}

	url, err := url.Parse(appSourceUrl)
	if err != nil {
		return fmt.Errorf("error: failed to parse app source url: %w", err)
	}

	opts := dagger.GhReleaseCreateOpts{
		GenerateNotes: true,
		Files:         m.BinFiles,
		VerifyTag:     true,
		Repo:          url.Path,
		Token:         token,
	}

	log.Printf("Publishing release for tag %s with %d files...\n", m.Git.Tag, len(m.BinFiles))

	//err = dag.Gh().Release().Create(ctx, m.Git.Tag, appName, opts)
	if err := dag.Gh().Release().Create(ctx, m.Git.Tag, m.Git.Tag, opts); err != nil {
		return fmt.Errorf("failed to create GitHub release: %w", err)
	}

	log.Println("Release published successfully!")

	return nil
}
