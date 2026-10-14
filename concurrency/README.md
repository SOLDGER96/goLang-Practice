## Concurrency in Go

By default, Go programs run sequentially—one operation must finish completely before the next one begins. Concurrency allows your program to deal with multiple tasks at the same time, making your code significantly faster and more efficient.

### 🧵 1. Goroutines
A goroutine is a lightweight thread managed by the Go runtime. Every Go program starts with one main goroutine (`main()`). To launch a concurrent background task, simply place the `go` keyword in front of a function call.

```go
package main

import "fmt"

func sayHello() {
    fmt.Println("Hello from the goroutine!")
}

func main() {
    // This launches concurrently
    go sayHello() 
    
    fmt.Println("Hello from main!")
    // Note: If main() finishes before the goroutine, the program exits and the goroutine is terminated.
}
```

### 🚦 2. WaitGroups (`sync.WaitGroup`)
Because the `main()` function does not wait for background goroutines to finish, we use a `WaitGroup` to synchronize them. It acts as a counter:
*   `Add(1)`: Increases the counter before starting a goroutine.
*   `Done()`: Decreases the counter when the goroutine finishes.
*   `Wait()`: Blocks the program from exiting until the counter reaches 0.

```go
package main

import (
    "fmt"
    "sync"
)

func sayHello(wg *sync.WaitGroup) {
    // Ensure Done() is called when the function finishes
    defer wg.Done() 
    fmt.Println("Hello from the goroutine!")
}

func main() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go sayHello(&wg)
    
    fmt.Println("Hello from main!")
    
    // Pauses main() until the counter hits 0
    wg.Wait() 
}
```

### 📬 3. Channels
Channels allow running goroutines to safely send and receive data to one another. Channels block by default, meaning if a function is waiting to receive data from a channel, it will pause until another goroutine sends the data. This provides built-in synchronization without needing a `WaitGroup`.

*   **Send Data:** `channel <- data`
*   **Receive Data:** `variable <- channel`

```go
package main

import "fmt"

func sendData(c chan string) {
    // Pushing data INTO the channel
    c <- "Hello through the channel!" 
}

func main() {
    // Creating a channel that strictly passes strings
    mychan := make(chan string)
    
    go sendData(mychan)
    
    // Pulling data OUT of the channel. 
    // main() blocks here until the data arrives.
    recData := <-mychan 
    fmt.Println(recData)
}
```

### 🛡️ 4. Mutexes (`sync.Mutex`)
When multiple goroutines try to read and modify the exact same variable at the exact same time, it creates a "race condition" that can corrupt your data. A Mutex (Mutual Exclusion lock) protects shared variables by locking them so only one goroutine can access them at a time.

Always use `defer` to unlock the Mutex. This guarantees the variable is unlocked even if the function crashes unexpectedly.

```go
package main

import (
    "fmt"
    "sync"
)

var balance int = 0
var mu sync.Mutex

func deposit(wg *sync.WaitGroup) {
    defer wg.Done()
    
    // Lock the door before making changes
    mu.Lock()
    // Guarantee it unlocks when the function exits
    defer mu.Unlock() 
    
    // Safe to modify the shared variable
    balance = balance + 100
}

func main() {
    var wg sync.WaitGroup
    
    // Launching 3 simultaneous deposits
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go deposit(&wg)
    }
    
    wg.Wait()
    fmt.Printf("Final Balance: $%d\n", balance) // Output: Final Balance: $300
}
```