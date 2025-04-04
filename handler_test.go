package slogger

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestContextHandler(t *testing.T) {
	tests := []struct {
		name          string
		setupOption   func() *Option
		setupContext  func() context.Context
		logFunc       func(ctx context.Context)
		expectedAttrs map[string]interface{}
	}{
		{
			name: "with_trace_id",
			setupOption: func() *Option {
				return NewOption().
					WithContextKey("trace-id")
			},
			setupContext: func() context.Context {
				ctx := context.Background()
				return context.WithValue(ctx, "trace-id", "abc-123")
			},
			logFunc: func(ctx context.Context) {
				slog.InfoContext(ctx, "test message")
			},
			expectedAttrs: map[string]interface{}{
				"trace-id": "abc-123",
			},
		},
		{
			name: "with_span_id",
			setupOption: func() *Option {
				return NewOption().
					WithSpanIdKey("span-id")
			},
			setupContext: func() context.Context {
				ctx := context.Background()
				spanId := 0
				return context.WithValue(ctx, "span-id", &spanId)
			},
			logFunc: func(ctx context.Context) {
				slog.InfoContext(ctx, "test message")
			},
			expectedAttrs: map[string]interface{}{
				"span-id": float64(1), // JSON 숫자는 float64로 파싱됨
			},
		},
		{
			name: "with_nil_context",
			setupOption: func() *Option {
				return NewOption()
			},
			setupContext: func() context.Context {
				return nil
			},
			logFunc: func(ctx context.Context) {
				slog.Info("message without context")
			},
			expectedAttrs: map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect log output to buffer
			var buf bytes.Buffer
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Initialize logger with test options
			option := tt.setupOption()
			err := Init(option)
			if err != nil {
				t.Fatalf("Init failed: %v", err)
			}

			// Get context and log
			ctx := tt.setupContext()
			tt.logFunc(ctx)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read logged output
			buf.ReadFrom(r)
			output := buf.String()

			// Parse last line if multiple lines
			lines := strings.Split(strings.TrimSpace(output), "\n")
			lastLine := lines[len(lines)-1]

			// Parse JSON
			var logData map[string]interface{}
			if err := json.Unmarshal([]byte(lastLine), &logData); err != nil {
				t.Fatalf("Failed to parse JSON: %v", err)
			}

			// Check expected attributes
			for k, v := range tt.expectedAttrs {
				if actual, ok := logData[k]; !ok {
					t.Errorf("Expected attribute %s not found", k)
				} else if actual != v {
					t.Errorf("Expected %s=%v, got %v", k, v, actual)
				}
			}
		})
	}
}

func TestWithGroups(t *testing.T) {
	// Redirect log output
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Initialize logger with trace ID context key
	err := Init(NewOption().WithContextKey("trace-id"))
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	// Create context with trace ID
	ctx := context.WithValue(context.Background(), "trace-id", "xyz-789")

	// Create a logger with a group and log a message
	logger := slog.Default().WithGroup("request")
	logger.InfoContext(ctx, "test message", "method", "GET")

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read output
	buf.ReadFrom(r)
	output := buf.String()

	t.Logf("Log output: %s", output)

	// Parse JSON
	var logData map[string]interface{}
	if err := json.Unmarshal([]byte(output), &logData); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Check request group was created
	requestGroup, ok := logData["request"].(map[string]interface{})
	if !ok {
		t.Error("request group not found in log")
		return
	}

	// Check method attribute in the request group
	method, ok := requestGroup["method"].(string)
	if !ok || method != "GET" {
		t.Errorf("Expected request.method=GET, got %v", requestGroup["method"])
	}

	// Check trace ID was preserved but is now inside the request group
	traceID, ok := requestGroup["trace-id"].(string)
	if !ok || traceID != "xyz-789" {
		t.Errorf("Expected request.trace-id=xyz-789, got %v", requestGroup["trace-id"])
	}
}
