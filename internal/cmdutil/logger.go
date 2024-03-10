package cmdutil

import (
	"log/slog"
	"os"
	"time"

	"gitlab.com/greyxor/slogor"
)

func SetLogger(debug bool) {
	logLevel := slog.LevelWarn
	if debug {
		logLevel = slog.LevelDebug
	}

	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
		TimeFormat: time.Stamp,
		Level:      logLevel,
		ShowSource: false,
	})))
}
