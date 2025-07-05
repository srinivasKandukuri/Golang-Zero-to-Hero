package main

// SOLID
/*

S -> single responsibility
	Class or module should have only one reason to change. a type should have single responsibility
	making the code easier to understand and maintain
	EX:User


O -> open and closed principle
	Software entities should be open for extension but closed for modification
	EX:  Shape interface => Area()

L -> Liskov substitution principle LSP
	The Objects of superclass should be replaceble with objects of subclasses without affecting the correctness of program.
	EX: Bird interface fly => makesound

I -> Interface Segregation principle
	Clients should not be forced to depend on interface they do not use
	This encourages creating smaller, more focused interfaces rather than large, monolethic ones
	EX: Document interface with read, write, print


D -> Dependency inversion principle
	High level module should not depending on the low level module
	EX : NotificationService and Emailservice , SMSservice

*/

// Single Responsiblity Principle
type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullname() string {
	return u.FirstName + u.LastName
}

type UserRepository struct {
	// dabase connection
}

func (r *UserRepository) Save(u *User) error {
	// save user to db
}

// Open closed principle
// open for extension closed for modification

type Shape interface {
	Area() float64
}

type Circle struct {
}

func (c Circle) Area() float64 {
	//
}

type Square struct {
}

func (s Square) Area() float64 {

}

// LSP
// Liskov subsitution principle
// super class should be replaceable with objects of  sub class without modifying the program

type Bird interface {
	MakeSound() string
}

type FlyingBird interface {
	Bird
	Fly() string
}

// ISP Interface segregration principle
type Read interface {
	read()
}
type Write interface {
	write()
}

type Copy interface {
	copy()
}

// Dependency inversion principle
//high level modules should not depending on low level modules

type MessageSender interface {
	send(to string, message string)
}

type EmailNotification struct{}

func (e *EmailNotification) send(to string, message string) {
	// send email
}

type SMSNotification struct{}

func (s *SMSNotification) send(to string, message string) {
	// send sms
}

// /////////////////////////////////////////////////////////////
type NotificationService struct {
	messageSender MessageSender
}

func NewNotificationService(messageSender MessageSender) *NotificationService {
	return &NotificationService{messageSender: messageSender}
}

func (n *NotificationService) Notify(to string, message string) {
	n.messageSender.send(to, message)
}

func main() {
	ap := NewNotificationService(EmailNotification{})
	ap.Notify("sk", "hello")
}
