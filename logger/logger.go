package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

var Level = new(slog.LevelVar)

func init() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      Level,
		TimeFormat: "15:04:05.000",
	})))
}
