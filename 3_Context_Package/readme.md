# 📖 Go Context Package

The `context` package in Go is a powerful and essential tool for managing deadlines, cancellations, and passing request-scoped values across API boundaries and between processes. It is particularly useful in concurrent programming, where you need to control the lifecycle of operations and prevent resource leakage.

## 🧠 What is the context package?

The `context` package allows you to:

- Manage timeouts, deadlines, and cancellations across goroutines.
- Handle HTTP requests, database calls, or long-running processes effectively.

### Why use the `context` package?

- ✅ Avoid leaving orphaned goroutines.
- ✅ Cancel operations if a client disconnects.
- ✅ Pass metadata like user authentication or request IDs.

## 📚 Core Functions in the context Package

### 1. `context.Background()`
- Returns an empty context.
- Ideal for the root context in `main()` functions or tests.

```go
ctx := context.Background()
```

### 2. `context.TODO()`
- Placeholder context when you are unsure about which context to use.
- Useful for future implementation.

```go
ctx := context.TODO()
```

### 3. `context.WithCancel()`
- Returns a derived context that can be canceled.
- Useful when you want to manually trigger cancellation.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

### 4. `context.WithDeadline()`
- Sets a deadline for the context.
- Automatically cancels the context when the deadline is reached.

```go
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
defer cancel()
```

### 5. `context.WithTimeout()`
- Similar to `WithDeadline` but sets a timeout duration instead of a fixed deadline.

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

### 6. `context.WithValue()`
- Attaches key-value pairs to the context.
- Useful for passing request-scoped data (e.g., user IDs, auth tokens).

```go
type key string

ctx := context.WithValue(context.Background(), key("userID"), 1234)
userID := ctx.Value(key("userID"))
```

## 📘 Additional Resources
- [Official Go Documentation - context package](https://pkg.go.dev/context)
- [Effective Go: Context Usage](https://golang.org/doc/effective_go.html#context)

---

Feel free to contribute or raise issues if you have any questions! 🚀

