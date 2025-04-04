package slogger

import (
	"log/slog"
	"os"
)

// Init initializes the logger with the provided options.
// If opt is nil, default options will be used.
// It sets up a JSON logger with appropriate time formatting and level.
// 컴파일 시 -trimpath 옵션을 추가하는 것을 권장합니다.
func Init(opt *Option) error {
	if opt != nil {
		option = *opt
	}

	var h contextHandler
	h.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     option.logLevel,
		AddSource: option.addSource, // 원본 소스 위치를 추가하는 옵션 적용
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
