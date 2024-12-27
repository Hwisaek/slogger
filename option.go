package slogger

import (
	"log/slog"
)

type Option struct {
	keyList    []string
	spanIdKey  *string
	timeFormat *string
	logLevel   *slog.Level
	addSource  bool
}

func NewOption() *Option {
	return &Option{
		timeFormat: &[]string{"2006-01-02T15:04:05.000-07:00"}[0],
		logLevel:   &[]slog.Level{slog.LevelDebug}[0],
	}
}

func (r Option) WithContextKey(key string) *Option {
	r.keyList = append(r.keyList, key)
	return &r
}

func (r Option) WithSpanIdKey(spanIdKey string) *Option {
	r.spanIdKey = &spanIdKey
	r.keyList = append(r.keyList, spanIdKey)
	return &r
}

func (r Option) WithTimeFormat(timeFormat string) *Option {
	r.timeFormat = &timeFormat
	return &r
}

func (r Option) WithLogLevel(logLevel slog.Level) *Option {
	r.logLevel = &logLevel
	return &r
}

func (r Option) WithAddSource(addSource bool) *Option {
	r.addSource = addSource
	return &r
}
