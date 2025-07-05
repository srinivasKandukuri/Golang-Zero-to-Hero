package main

import "fmt"

/*
✅ 3. Event Systems
Interfaces enable dynamic event handling where you can register and execute handlers without knowing their implementation.

Why use interfaces in event systems?

Loose coupling: Add/remove handlers without modifying existing code.
Extensibility: Easily introduce new event types
*/

type Event interface {
	Handle(eventData string)
}

type EventDispatcher struct {
	handlers []Event
}

func (e *EventDispatcher) Register(event Event) {
	e.handlers = append(e.handlers, event)
}

func (e *EventDispatcher) Dispatch(eventData string) {
	for _, v := range e.handlers {
		v.Handle(eventData)
	}
}

type EmailNotifier struct{}

func (e EmailNotifier) Handle(eventData string) {
	fmt.Println("sent email", eventData)
}

type SMSNotifier struct{}

func (e SMSNotifier) Handle(eventData string) {
	fmt.Println("sent sms", eventData)
}

func main() {
	e := EventDispatcher{}
	e.Register(EmailNotifier{})
	e.Register(SMSNotifier{})
	e.Dispatch("user sign up completed")
}
