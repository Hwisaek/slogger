package slogger

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
)

type contextHandler struct {
	slog.Handler
	opt *Option
}

func (h contextHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(h.observe(ctx)...)
	return h.Handler.Handle(ctx, r)
}

func (h contextHandler) observe(ctx context.Context) (as []slog.Attr) {
	option := h.opt
	if option == nil {
		return
	}

	if option.addSource {
		_, file, line, _ := runtime.Caller(4)
		codePath := fmt.Sprintf("%s:%d", file, line)

		as = append(as, slog.Attr{
			Key:   slog.SourceKey,
			Value: slog.StringValue(codePath),
		})
	}

	for _, key := range option.keyList {
		v := ctx.Value(key)

		if option.spanIdKey != nil {
			if key == *option.spanIdKey {
				if order, ok := v.(*int); ok {
					*order++
				} else {
					zero := 0
					v = &zero
				}
			}
		}

		as = append(as, slog.Attr{
			Key:   key,
			Value: slog.AnyValue(v),
		})
	}

	return
}
