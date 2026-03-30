package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	logFile     *os.File
	logMutex    sync.Mutex
	logFilePath string
)

// ErrorEntry represents a logged error
type ErrorEntry struct {
	Timestamp   time.Time `json:"timestamp"`
	Component   string    `json:"component"`
	Function    string    `json:"function"`
	Error       string    `json:"error"`
	Context     string    `json:"context,omitempty"`
	Severity    string    `json:"severity"` // ERROR, WARNING, INFO
	Resolved    bool      `json:"resolved"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
}

// InitLogger initializes the error logging system
func InitLogger(dataDir string) error {
	logMutex.Lock()
	defer logMutex.Unlock()

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	logFilePath = filepath.Join(dataDir, "errors.log")

	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	return nil
}

// Log records an error to the log file
func Log(component, function, errMsg, context string, severity string) {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFile == nil {
		fmt.Fprintf(os.Stderr, "Logger not initialized: %s\n", errMsg)
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] [%s] [%s/%s] %s", timestamp, severity, component, function, errMsg)

	if context != "" {
		logEntry += fmt.Sprintf(" | Context: %s", context)
	}

	logEntry += "\n"

	// Write to file
	if _, err := logFile.WriteString(logEntry); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log file: %v\n", err)
	}

	// Also write to stderr for ERROR severity
	if severity == "ERROR" {
		fmt.Fprintf(os.Stderr, logEntry)
	}
}

// Error logs an ERROR severity message
func Error(component, function, errMsg string, context ...string) {
	ctx := ""
	if len(context) > 0 {
		ctx = context[0]
	}
	Log(component, function, errMsg, ctx, "ERROR")
}

// Warning logs a WARNING severity message
func Warning(component, function, warnMsg string, context ...string) {
	ctx := ""
	if len(context) > 0 {
		ctx = context[0]
	}
	Log(component, function, warnMsg, ctx, "WARNING")
}

// Info logs an INFO severity message
func Info(component, function, infoMsg string, context ...string) {
	ctx := ""
	if len(context) > 0 {
		ctx = context[0]
	}
	Log(component, function, infoMsg, ctx, "INFO")
}

// CloseLogger closes the log file
func CloseLogger() error {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFile != nil {
		return logFile.Close()
	}
	return nil
}

// GetLogFilePath returns the path to the log file
func GetLogFilePath() string {
	return logFilePath
}

// ReadErrors reads all errors from the log file
func ReadErrors() ([]string, error) {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFilePath == "" {
		return nil, fmt.Errorf("logger not initialized")
	}

	content, err := os.ReadFile(logFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	lines := string(content)
	if lines == "" {
		return []string{}, nil
	}

	// Split by newlines
	var errors []string
	for _, line := range splitLines(lines) {
		if line != "" {
			errors = append(errors, line)
		}
	}

	return errors, nil
}

// ClearErrors clears the error log (use after resolving errors)
func ClearErrors() error {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFilePath == "" {
		return fmt.Errorf("logger not initialized")
	}

	// Truncate the file
	return os.Truncate(logFilePath, 0)
}

// splitLines splits a string by newlines
func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, c := range s {
		if c == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(c)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
