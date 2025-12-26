package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var levelNames = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

// Logger handles logging for the interpreter
type Logger struct {
	level     LogLevel
	output    io.Writer
	traceMode bool
	indentLvl int
}

var defaultLogger *Logger

func init() {
	defaultLogger = &Logger{
		level:     INFO,
		output:    os.Stdout,
		traceMode: false,
		indentLvl: 0,
	}
}

// New creates a new logger instance
func New(level LogLevel, output io.Writer) *Logger {
	return &Logger{
		level:     level,
		output:    output,
		traceMode: false,
		indentLvl: 0,
	}
}

// SetLevel sets the minimum log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// SetTraceMode enables or disables trace mode
func (l *Logger) SetTraceMode(enabled bool) {
	l.traceMode = enabled
}

// EnableTrace enables trace execution mode
func EnableTrace() {
	defaultLogger.SetTraceMode(true)
}

// DisableTrace disables trace execution mode
func DisableTrace() {
	defaultLogger.SetTraceMode(false)
}

// IsTraceEnabled returns true if trace mode is enabled
func IsTraceEnabled() bool {
	return defaultLogger.traceMode
}

// log writes a log message if it meets the minimum level
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	timestamp := time.Now().Format("15:04:05.000")
	levelStr := levelNames[level]
	message := fmt.Sprintf(format, args...)

	fmt.Fprintf(l.output, "[%s] %s: %s\n", timestamp, levelStr, message)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Trace logs execution trace information
func (l *Logger) Trace(format string, args ...interface{}) {
	if !l.traceMode {
		return
	}

	indent := ""
	for i := 0; i < l.indentLvl; i++ {
		indent += "  "
	}

	message := fmt.Sprintf(format, args...)
	fmt.Fprintf(l.output, "[TRACE] %s%s\n", indent, message)
}

// IncreaseIndent increases trace indentation
func (l *Logger) IncreaseIndent() {
	l.indentLvl++
}

// DecreaseIndent decreases trace indentation
func (l *Logger) DecreaseIndent() {
	if l.indentLvl > 0 {
		l.indentLvl--
	}
}

// Package-level convenience functions using default logger
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

func Trace(format string, args ...interface{}) {
	defaultLogger.Trace(format, args...)
}

func IncreaseIndent() {
	defaultLogger.IncreaseIndent()
}

func DecreaseIndent() {
	defaultLogger.DecreaseIndent()
}

func SetLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}
