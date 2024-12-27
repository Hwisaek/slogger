package slogger

import (
	"log/slog"
	"os"
)

// Init
// 컴파일 시 -trimpath 옵션을 추가하는 것을 권장
func Init(opt *Option) error {
	var h contextHandler
	h.opt = opt
	h.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: opt.logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				if opt.timeFormat != nil {
					a.Value = slog.StringValue(a.Value.Time().Format(*opt.timeFormat))
				}
			}

			return a
		},
	})

	slog.SetDefault(slog.New(h))
	return nil
}
