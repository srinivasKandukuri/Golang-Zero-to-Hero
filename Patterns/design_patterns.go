package main

import (
	"fmt"
	"sync"
	"time"
)

/*

Explanation of Patterns and Examples

Singleton: Ensures one instance of a database connection, common in applications needing a single DB pool.

Factory: Creates logger instances based on configuration, used in logging systems to switch between console, file, or cloud logging.

Builder: Builds server configurations step-by-step, ideal for setting up HTTP servers with flexible parameters.

Adapter: Bridges a legacy printer to a modern interface, useful when integrating old systems with new APIs.

Observer: Notifies multiple services (e.g., email, SMS) of events, common in event-driven systems like order processing.

Strategy: Switches caching mechanisms (in-memory vs. Redis), used in systems requiring flexible storage backends.

Decorator: Adds logging and metrics to HTTP handlers, widely used in middleware for web servers.

Command: Queues server start commands, useful in task scheduling or job queue systems.

Worker Pool: Processes jobs concurrently with a fixed number of workers, common in task-heavy systems like image processing or data pipelines.

These patterns leverage Go’s interfaces, structs, and concurrency features, aligning with its idiomatic style. Run the code to see each pattern in action. Let me know if you need deeper explanations or additional patterns!
*/

// 1. Singleton Pattern
// Ensures a single instance of a struct, useful for shared resources like database connections.
type SingletonDB struct {
	conn string
}

var singletonInstance *SingletonDB
var singletonMutex sync.Mutex

func GetDBInstance() *SingletonDB {
	singletonMutex.Lock()
	defer singletonMutex.Unlock()
	if singletonInstance == nil {
		singletonInstance = &SingletonDB{conn: "connected to postgres"}
	}
	return singletonInstance
}

// Real-world example: Database connection pool
func ExampleSingleton() {
	db1 := GetDBInstance()
	db2 := GetDBInstance()
	fmt.Printf("DB1: %p, DB2: %p, Same instance: %v\n", db1, db2, db1 == db2)
}

// 2. Factory Pattern
// Creates objects without specifying the exact type, useful for extensible systems.
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}
type FileLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

func (f *FileLogger) Log(message string) {
	fmt.Println("File:", message)
}

func LoggerFactory(logType string) Logger {
	switch logType {
	case "console":
		return &ConsoleLogger{}
	case "file":
		return &FileLogger{}
	default:
		return &ConsoleLogger{}
	}
}

// Real-world example: Logging system configuration
func ExampleFactory() {
	logger := LoggerFactory("console")
	logger.Log("Hello, World!")
	logger = LoggerFactory("file")
	logger.Log("Hello, World!")
}

// 3. Builder Pattern
// Constructs complex objects step-by-step, ideal for configurable setups.
type ServerConfig struct {
	Host string
	Port int
	SSL  bool
}

type ServerConfigBuilder struct {
	config ServerConfig
}

func NewServerConfigBuilder() *ServerConfigBuilder {
	return &ServerConfigBuilder{config: ServerConfig{Port: 8080}} // Default port
}

func (b *ServerConfigBuilder) SetHost(host string) *ServerConfigBuilder {
	b.config.Host = host
	return b
}

func (b *ServerConfigBuilder) SetPort(port int) *ServerConfigBuilder {
	b.config.Port = port
	return b
}

func (b *ServerConfigBuilder) EnableSSL(enable bool) *ServerConfigBuilder {
	b.config.SSL = enable
	return b
}

func (b *ServerConfigBuilder) Build() ServerConfig {
	return b.config
}

// Real-world example: HTTP server configuration
func ExampleBuilder() {
	config := NewServerConfigBuilder().
		SetHost("localhost").
		SetPort(443).
		EnableSSL(true).
		Build()
	fmt.Printf("Server Config: %+v\n", config)
}

// 4. Adapter Pattern
// Allows incompatible interfaces to work together, common in legacy integrations.
type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintLegacy(message string) {
	fmt.Println("Legacy Printer:", message)
}

type ModernPrinter interface {
	Print(message string)
}

type PrinterAdapter struct {
	legacyPrinter *LegacyPrinter
}

func (pa *PrinterAdapter) Print(message string) {
	pa.legacyPrinter.PrintLegacy(message)
}

// Real-world example: Adapting old logging system to new interface
func ExampleAdapter() {
	legacyPrinter := &LegacyPrinter{}
	adapter := &PrinterAdapter{legacyPrinter: legacyPrinter}
	adapter.Print("Adapted message")
}

// 5. Observer Pattern
// Notifies multiple objects of state changes, leveraging Go's concurrency.
type Subject struct {
	observers []Observer
}

type Observer interface {
	Update(data string)
}

type EmailObserver struct{}

func (eo *EmailObserver) Update(data string) {
	fmt.Println("Email notification:", data)
}

type SMSObserver struct{}

func (so *SMSObserver) Update(data string) {
	fmt.Println("SMS notification:", data)
}

func (s *Subject) AddObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(data string) {
	for _, observer := range s.observers {
		observer.Update(data)
	}
}

// Real-world example: Event notification system
func ExampleObserver() {
	subject := &Subject{}
	subject.AddObserver(&EmailObserver{})
	subject.AddObserver(&SMSObserver{})
	subject.Notify("Order placed!")
}

// 6. Strategy Pattern
// Defines interchangeable algorithms, useful for runtime behavior swapping.
type CacheStrategy interface {
	Store(key, value string)
}

type InMemoryCache struct{}

func (im *InMemoryCache) Store(key, value string) {
	fmt.Println("Storing in memory:", key, value)
}

type RedisCache struct{}

func (rc *RedisCache) Store(key, value string) {
	fmt.Println("Storing in Redis:", key, value)
}

type Cache struct {
	strategy CacheStrategy
}

func (c *Cache) SetStrategy(strategy CacheStrategy) {
	c.strategy = strategy
}

func (c *Cache) Store(key, value string) {
	c.strategy.Store(key, value)
}

// Real-world example: Caching system
func ExampleStrategy() {
	cache := &Cache{}
	cache.SetStrategy(&InMemoryCache{})
	cache.Store("user:1", "Alice")
	cache.SetStrategy(&RedisCache{})
	cache.Store("user:1", "Bob")
}

// 7. Decorator Pattern
// Extends functionality dynamically, often used in middleware.
type HandlerFunc func(string)

func LoggingDecorator(handler HandlerFunc) HandlerFunc {
	return func(s string) {
		fmt.Println("Logging:", s)
		handler(s)
	}
}

func MetricsDecorator(handler HandlerFunc) HandlerFunc {
	return func(s string) {
		fmt.Println("Metrics recorded for:", s)
		handler(s)
	}
}

// Real-world example: HTTP middleware
func ExampleDecorator() {
	handler := func(s string) { fmt.Println("Handling:", s) }
	decorated := LoggingDecorator(MetricsDecorator(handler))
	decorated("Request received")
}

// 8. Command Pattern
// Encapsulates requests as objects, useful for task queues.
type Command interface {
	Execute()
}

type StartServerCommand struct {
	server string
}

func (c *StartServerCommand) Execute() {
	fmt.Println("Starting server:", c.server)
}

type TaskQueue struct {
	commands []Command
}

func (tq *TaskQueue) AddCommand(c Command) {
	tq.commands = append(tq.commands, c)
}

func (tq *TaskQueue) Process() {
	for _, cmd := range tq.commands {
		cmd.Execute()
	}
}

// Real-world example: Job queue for server management
func ExampleCommand() {
	queue := &TaskQueue{}
	queue.AddCommand(&StartServerCommand{server: "web1"})
	queue.AddCommand(&StartServerCommand{server: "web2"})
	queue.Process()
}

// 9. Concurrency Pattern: Worker Pool
// Manages a pool of workers for parallel task processing.
type Job struct {
	id int
}

func Worker(id int, jobs <-chan Job, results chan<- string) {
	for job := range jobs {
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- fmt.Sprintf("Worker %d processed job %d", id, job.id)
	}
}

func ExampleWorkerPool() {
	const numJobs = 10
	const numWorkers = 3
	jobs := make(chan Job, numJobs)
	results := make(chan string, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go Worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{id: j}
	}
	close(jobs)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}

func main() {
	fmt.Println("Singleton Example:")
	ExampleSingleton()
	fmt.Println("\nFactory Example:")
	ExampleFactory()
	fmt.Println("\nBuilder Example:")
	ExampleBuilder()
	fmt.Println("\nAdapter Example:")
	ExampleAdapter()
	fmt.Println("\nObserver Example:")
	ExampleObserver()
	fmt.Println("\nStrategy Example:")
	ExampleStrategy()
	fmt.Println("\nDecorator Example:")
	ExampleDecorator()
	fmt.Println("\nCommand Example:")
	ExampleCommand()
	fmt.Println("\nWorker Pool Example:")
	ExampleWorkerPool()
}
