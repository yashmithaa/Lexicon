package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLoggerCreation(t *testing.T) {
	var buf bytes.Buffer
	logger := New(DEBUG, &buf)

	if logger == nil {
		t.Fatal("New() returned nil")
	}
	if logger.level != DEBUG {
		t.Errorf("logger.level = %v, want DEBUG", logger.level)
	}
	if logger.traceMode {
		t.Error("logger.traceMode should be false by default")
	}
}

func TestLogLevels(t *testing.T) {
	tests := []struct {
		level    LogLevel
		name     string
		logFunc  func(*Logger, string, ...interface{})
		expected string
	}{
		{DEBUG, "DEBUG", (*Logger).Debug, "DEBUG"},
		{INFO, "INFO", (*Logger).Info, "INFO"},
		{WARN, "WARN", (*Logger).Warn, "WARN"},
		{ERROR, "ERROR", (*Logger).Error, "ERROR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := New(DEBUG, &buf)
			tt.logFunc(logger, "test message")
			output := buf.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("log output should contain %q, got: %q", tt.expected, output)
			}
		})
	}
}

func TestTraceMode(t *testing.T) {
	var buf bytes.Buffer
	logger := New(DEBUG, &buf)

	logger.Trace("trace message 1")
	if buf.Len() > 0 {
		t.Error("Trace should not log when traceMode is false")
	}

	logger.SetTraceMode(true)
	logger.Trace("trace message 2")
	output := buf.String()
	if !strings.Contains(output, "[TRACE]") {
		t.Error("Trace output should contain [TRACE] prefix")
	}
}

func TestGlobalFunctions(t *testing.T) {
	EnableTrace()
	if !IsTraceEnabled() {
		t.Error("IsTraceEnabled() should return true after EnableTrace()")
	}
	DisableTrace()
	if IsTraceEnabled() {
		t.Error("IsTraceEnabled() should return false after DisableTrace()")
	}
}
