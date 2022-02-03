# Notes From Lectures


## go routine

- A goroutine exits when its code is complete
- When the main goroutine is complete, all other goroutines exit
- A goroutine may not complete its execution because main completes early

```go
 func main() {
     go fmt.Printf("New routine")
     fmt.Printf("Main routine")
}
```


##  Sync WaitGroup

```go
func foo(wg *sync.WaitGroup) {
    fmt.Printf("New routine")
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go foo(&wg)
    wg.Wait()
    
    fmt.Printf("Main routine")
}
```


## Channels

- Channel communication is synchronous
- Blocking is the same as waiting for communication
- Receiving and ignoring the result is same as a `Wait()`

```go
// Create a channel
c := make(chan int)
// Send data on a channel
c <- 3
// Receive data from a channel
x := <- c
```

```go
func prod(v1 int, v2 int, c chan int) {
    c <- v1 * v2
}

func main() {
    c := make(chan int)
    go prod(1, 2, c)
    go prod(3, 4, c)
    a := <- c
    b := <- c
    fmt.Println(a * b)
}
```


## Buffered Channels

- Channels can contain a limited number of objects
- Default size 0 (unbuffered)
- Capacity is the number of objects it can hold in transit
- Optional argument to `make()` defines channel capacity
    - `c := make(chan int, 3)`
- Sending only blocks if **buffer is full**
- Receiving only blocks if **buffer is empty**
 
**Channel Blocking, Receive**

```go
/*
Channel with capacity 1
One value sent to buffer
Two values are tried to be read
*/
// Task 1
// First receive blocks until send occurs
c <- 3
// Task 2
// Second receive blocks forever
a := <- c
b := <- c
```

**Channel Blocking, Receive**

```go
/*
Channel with capacity 1
Two values sent to buffer
One value is tried to be read
*/
// Task 1
// Second send blocks until receive is done
c <- 3
c <- 4
// Task 2
// Receive can block until first send is done
a := <- c
```
