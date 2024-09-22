package logger

import (
	"log/slog"
	"os"
)

func NewText() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}

func NewJson() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, nil))
}
