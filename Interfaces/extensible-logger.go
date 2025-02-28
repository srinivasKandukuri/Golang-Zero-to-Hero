package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("[console]:", message)
}

type FileLogger struct{}

func (f FileLogger) Log(message string) {
	fmt.Println("[file]: ", message)
}

type CloudLogger struct{}

func (c CloudLogger) Log(message string) {
	fmt.Println("[cloud]:", message)
}

func LogMessage(l Logger, message string) {
	l.Log(message)
}

func main() {
	LogMessage(ConsoleLogger{}, "system started")
	LogMessage(FileLogger{}, "file saved")
	LogMessage(CloudLogger{}, "cloud event triggered")
}

/*
✅ Benefits:
Flexible: Log to any system by implementing the Logger interface.
Extensible: Easily add new log targets (e.g., Kafka, Elasticsearch).
Isolated Logic: LogMessage doesn't care how logs are processed.
*/
