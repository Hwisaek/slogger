package tt

import (
	"github.com/Hwisaek/slogger"
	"log/slog"
	"testing"
)

func TestWithOutContext(t *testing.T) {
	option := slogger.NewOption().
		WithAddSource(true)
	if err := slogger.Init(option); err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info("info")
	slog.Debug("debug")
	slog.Warn("warn")
	slog.Error("error")
}
