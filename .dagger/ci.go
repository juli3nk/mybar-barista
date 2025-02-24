package main

import (
	"context"

	"dagger/mybar-barista/internal/dagger"
)

// Lint commit messages
func (m *MybarBarista) LintCommitMsg(
	ctx context.Context,
	args []string,
) (string, error) {
	return dag.Commitlint().
		Lint(m.Worktree, dagger.CommitlintLintOpts{Args: args}).
		Stdout(ctx)
}
