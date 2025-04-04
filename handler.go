package slogger

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
)

// contextHandler implements slog.Handler with support for context values
type contextHandler struct {
	slog.Handler
}

// Handle handles a log record and adds context values to the record
func (h contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if ctx != nil {
		// Only extract context values if context is not nil
		r.AddAttrs(h.observe(ctx)...)
	}
	return h.Handler.Handle(ctx, r)
}

// WithAttrs returns a new handler with attributes added to the set
func (h contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return contextHandler{Handler: h.Handler.WithAttrs(attrs)}
}

// WithGroup returns a new handler with the specified group added
func (h contextHandler) WithGroup(name string) slog.Handler {
	return contextHandler{Handler: h.Handler.WithGroup(name)}
}

// observe extracts values from context based on configured keys
func (h contextHandler) observe(ctx context.Context) (as []slog.Attr) {
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
