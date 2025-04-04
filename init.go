package slogger

import (
	"log/slog"
	"os"
)

// Init initializes the logger with the provided options.
// If opt is nil, default options will be used.
// It sets up a JSON logger with appropriate time formatting and level.
// It is recommended to compile with the -trimpath option.
func Init(opt *Option) error {
	if opt != nil {
		option = *opt
	}

	var h contextHandler
	h.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     option.logLevel,
		AddSource: option.addSource, // Apply the option to add source location
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				a.Value = slog.StringValue(a.Value.Time().Format(option.timeFormat))
			}

			return a
		},
	})

	slog.SetDefault(slog.New(h))
	return nil
}
