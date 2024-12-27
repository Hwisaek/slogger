package tt

import (
	"context"
	"github.com/Hwisaek/slogger/v2"
	"log/slog"
	"testing"
)

func TestWithContext(t *testing.T) {
	option := slogger.NewOption().
		WithAddSource(true).
		WithSpanIdKey("test").
		WithContextKey("traceId")
	if err := slogger.Init(option); err != nil {
		slog.Error(err.Error())
		return
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "traceId", "123")
	ctx = context.WithValue(ctx, "test", &[]int{0}[0])

	slog.InfoContext(ctx, "info context")
	slog.DebugContext(ctx, "debug context")
	slog.WarnContext(ctx, "warn context")
	slog.ErrorContext(ctx, "error context")

}
