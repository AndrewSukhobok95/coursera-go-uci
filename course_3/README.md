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

- Common to iteratively read from a channel
- Iterates when sender calls `close(c)`

```go
for i := range c {
    fmt.Println(i)
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


## Select Statement

```go
select {
    case a = <- c1:
        fmt.Println(a)
    case b = <- c2:
        fmt.Println(b)
    case a = <- inchan:
        fmt.Println("Received a")
    case outchan <- b:
        fmt.Println("Sent b")
    case <- abort:
         return
    default:
        fmt.Println("default case to avoid blocking")
}
```

## Mutual Exclusion

- `Lock()` method puts the flag up – Shared variable in use
- If lock is already taken by a goroutine, `Lock()` blocks until the flag is put down
- `Unlock()` method puts the flag down – Done using shared variable
- When `Unlock()` is called, a blocked `Lock()` can proceed

```go
var i int = 0
var wg sync.WaitGroup
var mut sync.Mutex

func inc() {
    mut.Lock()
    i =i + 1
    mut.Unlock()
}

func main() {
    wg.Add(2)
    go inc()
    go inc()
    wg.Wait()
    fmt.Println(i)
}
```


## Synchronous Initialization

```go
var wg sync.WaitGroup
var on sync.Once

func setup() {
    fmt.Println("Init")
}

func dostuff() {
    on.Do(setup)
    fmt.Println("hello")
    wg.Done()
}

func main() {
    wg.Add(2)
    go dostuff()
    go dostuff()
    wg.Wait()
}
```



