package slogger

import (
	"log/slog"
	"testing"
)

func TestGetLogLevel(t *testing.T) {
	type args struct {
		logLevel slog.Level
	}
	tests := []struct {
		name string
		args args
		want slog.Level
	}{
		{
			name: "",
			args: args{
				logLevel: slog.LevelDebug,
			},
			want: slog.LevelDebug,
		},
		{
			name: "",
			args: args{
				logLevel: slog.LevelInfo,
			},
			want: slog.LevelInfo,
		},
		{
			name: "",
			args: args{
				logLevel: slog.LevelWarn,
			},
			want: slog.LevelWarn,
		},
		{
			name: "",
			args: args{
				logLevel: slog.LevelError,
			},
			want: slog.LevelError,
		},
		{
			name: "",
			args: args{},
			want: slog.LevelInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := NewOption().
				WithAddSource(false).
				WithLogLevel(tt.args.logLevel)
			if err := Init(option); err != nil {
				t.Error("init error")
			}

			if got := GetLogLevel(); got != tt.want {
				t.Errorf("GetLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
