# Core Go Concepts

## 1. What are Goroutines, and how are they different from threads?

**Goroutines** are lightweight, managed threads provided by Go’s runtime. You start a goroutine by prefixing a function
call with the `go` keyword.

**Differences from threads:**

- Goroutines are much lighter than OS threads; thousands can run concurrently in the same address space.
- The Go scheduler multiplexes many goroutines onto a smaller number of OS threads, handling their scheduling and
  execution.
- Goroutines have smaller initial stack sizes (typically 2 KB) that grow and shrink as needed, unlike fixed-size thread
  stacks.
- Communication and synchronization between goroutines are typically done using **channels**, not shared memory and
  locks.

---

## 2. Explain the concept of channels in Go. When and why would you use them?

**Channels** are typed pipes that allow goroutines to communicate and synchronize execution by sending and receiving
values.

**Syntax:**

```go
ch := make(chan int)
```

text

**Usage:**

- Used for safe data exchange between goroutines, avoiding explicit locks.
- Useful for implementing producer-consumer patterns, pipelines, and coordinating concurrent tasks.
- Channels can be buffered or unbuffered, affecting whether send/receive operations block.

**When to use:**  
When you need to coordinate work or share data between concurrent goroutines safely and idiomatically.

---

## 3. Describe the use of the defer statement in Go with an example.

The `defer` keyword schedules a function call to be executed after the surrounding function returns—whether by normal
return or panic.

**Common use:** Resource cleanup (e.g., closing files, unlocking mutexes).

**Example:**

```go
func main() {
defer fmt.Println("World")
fmt.Println("Hello")
}
// Output:
// Hello
// World
```

text
Deferred calls are executed in last-in, first-out (LIFO) order and even run if the function panics.

---

## 4. What is the difference between a slice and an array in Go?

| Feature     | Array                                   | Slice                          |
|-------------|-----------------------------------------|--------------------------------|
| Size        | Fixed at compile time                   | Dynamic, can grow/shrink       |
| Declaration | `[5]int{1,2,3,4,5}`                     | `[]int{1,2,3,4,5}`             |
| Underlying  | Stores the actual data                  | References an underlying array |
| Use case    | Rare, mostly for fixed-size collections | Common, flexible, idiomatic    |

---

## 5. How does Go handle memory management and garbage collection?

- Go uses **automatic garbage collection**: memory allocation and deallocation are managed by the Go runtime.
- The garbage collector is concurrent and optimized for low latency, reclaiming memory that is no longer referenced.
- Developers do not need to explicitly free memory, reducing the risk of leaks and dangling pointers.

---

## 6. What is the role of the init function in Go?

- The `init` function is a special function that runs automatically before the `main` function and after all
  package-level variables are initialized.
- Used for setup tasks such as initializing state, registering types, or validating environment/configuration.
- Each package can have multiple `init` functions (one per file); their execution order is file order within the
  package.

---

## 7. How do you handle errors in Go? Can you provide an example?

- Go uses **explicit error handling** with the built-in `error` interface.
- Functions that can fail return an error as their last return value.

**Idiomatic pattern:**

```go
f, err := os.Open("file.txt")
if err != nil {
// handle error
return err
}
defer f.Close()
// use f
```

text
This approach makes error handling explicit and visible, improving code reliability.

---

## 8. What is a Go interface, and how do you implement it? Why are interfaces important in Go?

- An **interface** is a type that specifies a set of method signatures. Any type that implements those methods satisfies
  the interface, implicitly—no explicit declaration needed.

**Example:**

```go
type Reader interface {
Read(p []byte) (n int, err error)
}

type MyReader struct{}
func (r MyReader) Read(p []byte) (n int, err error) { /.../ }
```

text
**Importance:** Enables polymorphism, decouples code, and supports testability and extensibility without inheritance.

---

## 9. Explain the concept of 'embedding' in Go structs.

- **Embedding** allows one struct type to include another struct type anonymously.
- Promotes code reuse and composition over inheritance.
- The embedded struct’s fields and methods become accessible on the outer struct.

**Example:**

```go
type Animal struct {
Name string
}
type Dog struct {
Animal
Breed string
}
// Dog has both Name and Breed fields.
```

text

---

## 10. What tools or libraries do you use for dependency management in Go?

- **Go modules** (introduced in Go 1.11, standard since 1.13) are the official dependency management system.
    - Commands: `go mod init`, `go mod tidy`, `go mod vendor`
- Older tools (now rarely used): `dep`, `govendor`, `glide`
- Go modules handle versioning, reproducible builds, and proxying dependencies.