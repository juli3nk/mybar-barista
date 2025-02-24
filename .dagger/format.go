package main

import (
	"context"
	"log"

	"dagger/mybar-barista/internal/dagger"

	"github.com/juli3nk/go-utils/filedir"
)

// Format Golang files
func (m *MybarBarista) FormatGo(
	ctx context.Context,
	checkOnlyModifiedFiles bool,
) (string, error) {
	opts := dagger.GoFmtOpts{}

	if checkOnlyModifiedFiles {
		files := filedir.FilterFileByExtension(m.Git.ModifiedFiles, "go")
		if len(files) == 0 {
			log.Println("No modified Go files to lint.")
			return "", nil
		}

		opts.Filedir = files
	}

	return dag.Go(goVersion, m.Worktree).Fmt(ctx, opts)
}
