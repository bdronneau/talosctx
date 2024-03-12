package cmdutil

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLogger(t *testing.T) {
	ctx := context.Background()
	assert.False(t, slog.Default().Enabled(ctx, slog.LevelDebug))

	SetLogger(true)

	assert.True(t, slog.Default().Enabled(ctx, slog.LevelDebug))
}
