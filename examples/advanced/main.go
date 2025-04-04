package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Hwisaek/slogger/v2"
)

func main() {
	// Initialize with advanced settings
	option := slogger.NewOption().
		WithAddSource(true).                       // Add source locations
		WithTimeFormat("2006-01-02 15:04:05.000"). // Set timestamp format
		WithLogLevel(slog.LevelDebug).             // Set log level to Debug
		WithSpanIdKey(slogger.ContextKeySpanId).   // Configure span ID key
		WithContextKey(slogger.ContextKeyTraceId)  // Add trace ID key

	if err := slogger.Init(option); err != nil {
		panic(err)
	}

	// Debug level logs will now be output
	slog.Debug("Debug log message - displayed as level is set to Debug")

	// HTTP server setup
	http.HandleFunc("/", handleRequest)
	slog.Info("Starting server...", "port", "8080", "environment", "development")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to start server", "error", err, "action", "retry_required")
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Create context for new request
	traceID := "trace-" + r.RemoteAddr
	ctx := context.WithValue(r.Context(), slogger.ContextKeyTraceId, traceID)
	spanID := 0
	ctx = context.WithValue(ctx, slogger.ContextKeySpanId, &spanID)

	// Log request
	slog.InfoContext(ctx, "Request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent(),
	)

	// Processing logic...
	slog.DebugContext(ctx, "Processing request started")

	// Check current span ID
	slog.InfoContext(ctx, "Current span ID", "span", spanID)

	// More processing...
	slog.DebugContext(ctx, "Processing request completed")

	// Response
	w.Write([]byte("Hello, world!"))

	// Log response
	slog.InfoContext(ctx, "Response sent", "status", 200)
}
