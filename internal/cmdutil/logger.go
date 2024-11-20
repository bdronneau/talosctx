package cmdutil

import (
	"log/slog"
	"os"
)

func SetLogger(debug bool) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if debug {
		opts.Level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
}
