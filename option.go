package slogger

import (
	"log/slog"
)

// Option represents configuration options for the logger
type Option struct {
	keyList    []string   // List of context keys to extract and include in logs
	spanIdKey  *string    // Optional key for span ID that will be auto-incremented
	timeFormat string     // Format for timestamp in logs
	logLevel   slog.Level // Minimum log level to output
	addSource  bool       // Whether to add source code location to logs
}

// Global default option
var option = *NewOption()

// NewOption creates a new Option instance with default values
func NewOption() *Option {
	return &Option{
		keyList:    []string{},
		timeFormat: "2006-01-02T15:04:05.000-07:00", // RFC3339 with milliseconds
		logLevel:   slog.LevelInfo,                  // Default to Info level
		addSource:  false,                           // Source location disabled by default
	}
}

// GetLogLevel returns the currently configured log level
func GetLogLevel() slog.Level {
	return option.logLevel
}

// WithContextKey adds a context key to extract and include in logs
func (r Option) WithContextKey(key string) *Option {
	r.keyList = append(r.keyList, key)
	return &r
}

// WithSpanIdKey configures a special span ID key that will be auto-incremented
// and included in logs. The key is also added to the context key list.
func (r Option) WithSpanIdKey(spanIdKey string) *Option {
	r.spanIdKey = &spanIdKey
	r.keyList = append(r.keyList, spanIdKey)
	return &r
}

// WithTimeFormat sets the format for timestamps in logs
// Uses Go's time format string (e.g., "2006-01-02T15:04:05.000-07:00")
func (r Option) WithTimeFormat(timeFormat string) *Option {
	r.timeFormat = timeFormat
	return &r
}

// WithLogLevel sets the minimum log level that will be output
// Logs below this level will be discarded
func (r Option) WithLogLevel(logLevel slog.Level) *Option {
	r.logLevel = logLevel
	return &r
}

// WithAddSource enables or disables adding source code location to logs
func (r Option) WithAddSource(addSource bool) *Option {
	r.addSource = addSource
	return &r
}
