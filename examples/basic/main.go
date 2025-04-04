package main

import (
	"context"
	"log/slog"

	"github.com/Hwisaek/slogger/v2"
)

func main() {
	// Initialize with default settings
	err := slogger.Init(nil)
	if err != nil {
		panic(err)
	}

	// Basic logging
	slog.Info("Basic info log")
	slog.Debug("Basic debug log - not displayed as default level is Info")
	slog.Warn("Warning log")
	slog.Error("Error log")

	// Structured logging
	slog.Info("Structured log", "user_id", "user123", "action", "login")

	// Context-based logging
	ctx := context.Background()
	slog.InfoContext(ctx, "Context log")

	// Using log groups
	logger := slog.Default().WithGroup("request")
	logger.Info("Message within log group", "method", "GET", "path", "/api/users")
}
