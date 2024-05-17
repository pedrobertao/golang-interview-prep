## 🤖  General questions 🤖

#### 1. What is Go? Explain its main features.

> **Answer:**
> Go, also known as Golang, is an open-source programming language developed by Google. It is statically typed, compiled, and known for its simplicity, efficiency, and concurrency support. Key features include garbage collection, rich standard library, goroutines for concurrency, and built-in tools for code formatting and testing.

#### 2. How do you declare a variable in Go?

> **Answer:**
> You can declare a variable using the var keyword or the short declaration operator :=. For example:

```
var x int
x = 10 // or
y := 20
```

#### 3. What are Goroutines and how do they differ from threads?

> **Answer:**
> Goroutines are lightweight threads managed by the Go runtime. They are more efficient than OS threads, allowing thousands of them to run concurrently. Goroutines use a multiplexing technique to achieve concurrency, whereas threads are managed by the OS and are heavier in terms of resource consumption.

#### 4. What is the purpose of the defer statement in Go?

> **Answer:**
> The defer statement is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. Deferred functions are executed in LIFO (Last In, First Out) order just before the surrounding function returns.

#### 5. Explain how Go handles errors.

> **Answer:**
> Go handles errors using a simple, explicit error type. Functions that can result in an error return an additional error value. Error handling is done by checking this returned value. The standard library provides the errors package to create and manipulate errors.

#### 6. What is a Go interface and how is it used?

> **Answer:**
> An interface in Go is a type that specifies a set of method signatures. A type implements an interface by implementing its methods. Interfaces are used to define behavior and achieve polymorphism. For example:
```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

#### 7. Describe Go's package management and how dependencies are handled.

> **Answer:**
> Go uses modules to manage dependencies, defined by a go.mod file. The go command provides various tools for managing dependencies, such as go get to download packages and go mod to manage modules and dependencies.

#### 8. What are channels in Go and how are they used?

> **Answer:**
> Channels are a typed conduit through which you can send and receive values with the channel operator, <-. They are used to synchronize goroutines and communicate data between them. Channels can be buffered or unbuffered.

#### 9. Explain how the garbage collector works in Go.

> **Answer:**
> Go uses a concurrent, mark-and-sweep garbage collector. It runs in the background, identifying and collecting unused memory objects while minimizing pauses in the execution of the program. The garbage collector uses multiple strategies to balance between memory efficiency and application latency.

#### 10. What are the best practices for writing concurrent programs in Go?

> **Answer:**
> *Best practices include:*
> - Avoid sharing memory by communicating through channels.
> - Use goroutines efficiently and avoid creating too many.
> - Properly handle synchronization using channels, mutexes, and other synchronization primitives.
> - Avoid data races by protecting shared data access.

#### 11. How does Go implement slices and what are the underlying mechanisms?

> **Answer:**
> Slices in Go are dynamically-sized, flexible views into the elements of an array. They are described by a struct containing a pointer to the array, the length, and the capacity. Operations on slices involve manipulation of these fields and may involve reallocation of underlying arrays when resizing.

#### 12. What is a select statement in Go and how is it used?

> **Answer:**
> The select statement allows a goroutine to wait on multiple communication operations. It blocks until one of its cases can proceed, then executes that case. It is used to handle multiple channels and avoid blocking on a single channel.

## ❗  Tricky Questions  ❗

#### 1. What is the output of the following Go program and why?

```
import "fmt"
package main
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

> **Answer:**
> hello
> world

> **Explanation:**
> The defer statement defers the execution of fmt.Println("world") until the surrounding function (main) returns, so "hello" is printed first, followed by "world".

### 2. How do you prevent a data race in the following Go code?
```
package main
import (
    "fmt"
    "time"
)
func main() {
    var x int
    go func() {
        x++
    }()
    time.Sleep(1 * time.Second)
    fmt.Println(x)
}
```

> **Answer:**
> To prevent a data race, you can use a synchronization primitive like a mutex or a channel. Here’s one way using a mutex:
```
package main
import (
    "fmt"
    "sync"
    "time"
)
func main() {
    var x int
    var mu sync.Mutex
    go func() {
        mu.Lock()
        x++
        mu.Unlock()
    }()
    time.Sleep(1 * time.Second)
    mu.Lock()
    fmt.Println(x)
    mu.Unlock()
}
```

#### 3. What will be the output of the following program? Explain why.
```
package main
import "fmt"
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
```

> **Answer:**
> [0 0 0 0 0 1 2 3]

> **Explanation:**
> The slice s is initially created with a length of 5, containing five zeroes. When append is called, it adds three more elements to the slice, resulting in a total length of 8.

#### 4. How does Go handle nil interfaces? What is a common pitfall associated with them?

> **Answer:**:
> A nil interface in Go is an interface value that holds neither a value nor a type. A common pitfall is checking only the value for nil, not considering that an interface can be non-nil while holding a nil concrete value. This can lead to unexpected behavior when comparing interfaces for nil.

#### 5. Describe the memory model of Go and its implications for concurrent programming.

> **Answer:**
> Go's memory model defines the behavior of memory access in concurrent programs. It ensures that read and write operations on variables are atomic and provides rules for the visibility of writes to other goroutines. This model requires proper synchronization to ensure safe concurrent access to shared variables, typically using channels or sync primitives.

#### 6. How would you implement a worker pool in Go?

> **Answer:**:
> A worker pool can be implemented using goroutines and channels. Here’s an example:
```
package main
import (
    "fmt"
    "sync"
)
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        results <- j * 2
    }
}
func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    var wg sync.WaitGroup
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    close(results)

    for r := range results {
        fmt.Println("Result:", r)
    }
}
```

#### 7. What are some common performance pitfalls in Go, and how can you avoid them?

> **Answer:**
> *Common performance pitfalls include:*
> - Excessive use of goroutines leading to high memory consumption.
> - Inefficient use of channels causing unnecessary blocking.
> - Large memory allocations and frequent garbage collection pauses.
> - Suboptimal use of the standard library functions.

> *These can be avoided by:*
> - Limiting the number of goroutines.
> - Using buffered channels appropriately.
> - Profiling the application to identify and optimize hot spots.
> - Leveraging efficient data structures and algorithms.

#### 8. Explain the difference between a pointer receiver and a value receiver in Go methods. When would you use each?

> **Answer:**
> A pointer receiver allows methods to modify the receiver's value and avoid copying it. It is used when the method needs to mutate the receiver or when the receiver is a large struct to avoid performance overhead of copying. A value receiver is used when the method does not need to modify the receiver, and copying is cheap.