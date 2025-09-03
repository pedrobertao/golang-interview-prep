# Concurrency and Parallelism in Golang

## 🧠 What is Concurrency?

**Concurrency** is about dealing with multiple tasks at the same time conceptually (handling many things at once, but not necessarily executing them simultaneously).  
In Go:

- Concurrency allows writing programs that can handle many tasks that are logically happening at the same time.
- Example: A web server handling thousands of connections by switching between them efficiently.

## ⚡ What is Parallelism?

**Parallelism** is about executing multiple tasks at exactly the same time (simultaneously on multiple CPU cores).  
In Go:

- Parallelism happens when goroutines are scheduled to run truly in parallel on different CPU cores.
- Example: Running CPU-bound calculations concurrently across cores.

---

## 🔹 Key difference:

| Concept     | Meaning                                                     | Example                                                                     |
| ----------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- |
| Concurrency | Managing multiple tasks, potentially interleaved.           | Handling thousands of HTTP requests with one thread, switching efficiently. |
| Parallelism | Performing multiple tasks simultaneously on multiple cores. | Processing chunks of a large array on separate CPU cores at the same time.  |

---

## 🏎️ How Golang Handles Concurrency and Parallelism

### Goroutines:

- **Goroutines** are lightweight threads managed by the Go runtime.
- `go f()` starts a new goroutine executing `f()` concurrently.

### Scheduler:

- Go uses a **work-stealing scheduler** to multiplex goroutines onto a smaller set of OS threads (`GOMAXPROCS` defines how many OS threads run in parallel).

### Channels:

- **Channels** enable communication between goroutines safely and efficiently.
- Example:
  ```go
  c := make(chan int)
  go func() {
      c <- 42
  }()
  fmt.Println(<-c)
  ```
