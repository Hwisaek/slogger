package slogger

import (
	"log/slog"
	"os"
)

// Init
// 컴파일 시 -trimpath 옵션을 추가하는 것을 권장
func Init(opt *Option) error {
	if opt != nil {
		option = *opt
	}

	var h contextHandler
	h.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: option.logLevel,
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
