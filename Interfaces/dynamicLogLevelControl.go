package main

import (
	"fmt"
	"sync"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
)

type DynamicLogger interface {
	Log(level LogLevel, message string)
	SetLevel(level LogLevel)
}

type CustomLogger struct {
	mu    sync.RWMutex
	level LogLevel
}

func (l *CustomLogger) Log(level LogLevel, message string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if level >= l.level {
		fmt.Printf("[%s] %s\n", l.levelToString(level), message)
	}
}

func (l *CustomLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Printf("[Logger] Switching log level to: %s\n", l.levelToString(level))
	l.level = level
}

// Helper: Convert LogLevel to string
func (l *CustomLogger) levelToString(level LogLevel) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	default:
		return "UNKNOWN"
	}
}

func NewCustomeLogger(level LogLevel) *CustomLogger {
	return &CustomLogger{level: level}
}

type App struct {
	logger DynamicLogger
}

// Run executes application tasks with logging
func (a *App) Run() {
	a.logger.Log(DebugLevel, "Starting the application...")
	a.logger.Log(InfoLevel, "Application is running")
}

func main() {

	logger := NewCustomeLogger(DebugLevel)

	app := App{logger: logger}
	// Run the application
	go app.Run()

	fmt.Println("Press Enter to switch to INFO level...")
	fmt.Scanln() // Wait for user input

	logger.SetLevel(InfoLevel)

	// Run again with updated log level
	app.Run()
}
