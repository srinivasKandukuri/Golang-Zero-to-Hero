package main

import "fmt"

type Zaplog interface {
	Log(message string)
}

type ConsoleLog struct{}

func (c ConsoleLog) Log(message string) {
	fmt.Println("[Console]:", message)
}

type FileLog struct{}

func (f FileLog) Log(message string) {
	fmt.Println("[File]:", message)
}

// Application using Logger interface
type Application struct {
	zapLogger Zaplog
}

func NewApplication(zapLogger Zaplog) *Application {
	return &Application{zapLogger: zapLogger}
}

func (a *Application) Run() {
	a.zapLogger.Log("Application started")
}

func main() {
	app := NewApplication(ConsoleLog{})

	app.Run()

	file := NewApplication(FileLog{})

	file.Run()
}

/*
1. Dependency Injection (DI) with Interfaces
Interfaces enable dependency injection, which allows swapping out implementations without changing core logic.

✅ Why is this useful?

Flexibility: Easily swap between logging mechanisms.
Testability: Mock the Logger interface in unit tests.
Separation of Concerns: Business logic doesn’t depend on how logs are stored.
*/
