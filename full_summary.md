# Notes From Lectures

## Basics

### Pointers

```go
var x int = 1
var y int
var ip *int    // ip is pointer to int

ip = &x        // ip now points to x
y = *ip        // y is now 1
```

```go
ptr := new(int)  // new() creates a variable and returns a pointer to it
*ptr = 3         // 3 is placed into address specified by ptr
```

### Variable Scope

- Place in the code where a variable can be accessed.

### Strings

```go
x := "Hi!"  // double quotes
```

**"unicode"** package

```go
IsDigit(r rune)
IsSpace(r rune)
IsLetter(r rune)
IsLower(r rune)
IsPunct(r rune)

ToUpper(r rune)
ToLower(r rune)
```

**"strings"** package

```go
Compare(a, b)
Contains(s, substr)
HasPrefix(s, prefix)
Index(s, substr)

Replace(s, old, new, n)
ToLower(s)
ToUpper(s)
TrimSpace(s)
```

**"strconv"** package

```go
Atoi(s)                             // ASCII to integer
Itoa(s)                             // int to ASCII string
FormatFloat(f, fmt, prec, bitSize)  // float to a string
ParseFloat(s, bitSize)              // string to a float
```

### Constants

```go
const x = 1.3
const (
    y = 4
    z = "Hi"
)

type Grades int
const (
    A Grades = iota
    B
    C
    D
    F
)
```

### Control Flow

```go
if <condition> {
    <consequent>
}

if x > 5 {
    fmt.Printf("Yup")
}

for <init>; <cond>; <update> {
    <stmts>
}

for i:=0; i<10; i++ {
    fmt.Printf("hi")
}

i = 0
for i<10 {
    fmt.Printf("hi")
    if i == 5 { continue } // skip iteration 5
    if i == 8 { break }    // break on 8th iteration
    i++
}

for {
    fmt.Printf("hi")
}


switch x {
case 1:
    fmt.Printf("case 1")
case 2:
    fmt.Printf("case 2")
default:
    fmt.Printf("default")
}

switch {
case x > 1:
    fmt.Printf("case 1")
case x < -1:
    fmt.Printf("case 2")
default:
    fmt.Printf("default")
}
```


### Scan

```go
var appleNum int

fmt.Printf("Number of apples?")

num, err := fmt.Scan(&appleNum)  // receives a pointer

fmt.Printf(appleNum)
```


### Arrays

```go
var x [5]int
x[0] = 2

var x [5]int = [5]{1,2,3,4,5}

x := [...]int{1,2,3,4}

x := [3]int{1,2,3}
for i, v range x {
    fmt.Printf("ind %d, val %d", i, v)
}
```


### Slices

```go
// 1ST WAY OF CREATING SLICE
arr := [...]string{"a", "b", "c", "d","e", "f", "g"}
s1 := arr[1:3]
s2 := arr[2:5]
fmt.Printf(len(s1), cap(s1))  // 2, 7

// 2ND WAY OF CREATING SLICE
sli := []int{1,2,3}            // slice, because no num.of elements or ...

// 3D WAY OF CREATING SLICE
sli := make([]int, 10)
sli := make([]int, 10, 15)     // type, length, capacity

// append element to a variable length slice
sli = append(sli, 100)
```


### Hash Tables

```go
var idMap map[string]int
idMap = make(map[string]int)

idMap := map[string]int {"joe": 123}

fmt.Printf(idMap["joe"])

idMap["jane"] = 465

delete(idMap, "joe")

id, p := idMap["joe"]    // id is a value, p is bool of presence of the key

fmt.Printf(len(idMap))

for key, val := range idMap {
    fmt.Printf(key, val)
}
```


### Structs

```go
type Person struct {
    name string
    addr string
    phone string
}

var p1 Person

p1.name = "joe"
x = p1.addr

p1 := new(Person)  // sets all fields to "zero" values

p1 := Person{name: "joe", addr: "a st.", phone: "123"}
```


### Protocol Packages

**"net/http"**

```go
http.Get("www.uci.edu")
```

**"net"**

```go
net.Dial("tcp", "uci.edu:80")
```


### JSON

**Marchalling**

```go
type struct Person {
    name string
    addr string
    phone string
}

p1 := Person(name: "Joe", addr: "a st.", phone: "123")
barr, err := json.Marshal(p1)   // barr - []byte representation of json

var p2 Person
err := json.Unmarchal(barr, &p2)
```


### File accessed

**"ioutil"** package

```go
// open and close are built in ioutil.ReadFile
// byte array dat and error e
dat, e := ioutil.ReadFile("test.txt")

dat = "Hello World!"
err := ioutil.WriteFile("outfile.txt", dat, 0777)
```

**"os"** package

```go
os.Open()
os.Close()
os.Read()
os.Write()

f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()

f, err := os.Create("outfile.txt")
bar := []byte{1,2,3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
```




## Functions

```go
func printX(x int) {
  fmt.Printf(x)
}

func foo(x int) int {
  return x + 1
}
y := foo(1)

func foo2(x int) (int, int) {
  return x, x + 1
}
a, b := foo2(3)
```

### Call by value / reference

- **call by value** is default in go:

```go
func foo(y int) {
    y = y+1
}
func main() {
    x := 2
    foo(x)
    fmt.Print(x)
}
```

- **call by reference** is possible by passing a pointer to the function:

```go
func foo(y *int) {
    *y = *y + 1
}
func main() {
    x := 2
    foo(&x)
    fmt.Print(x)
}
```

- **call by reference** for arrays:
    - Not go-way of passing an array
    - Better passing a slice

```go
// Passing an aray by reference
func foo(x [3]int) int {
    return x[0]
}
func main() {
    a := [3]int{1, 2, 3}
    fmt.Print(foo(a))
}

// Passing a slice
func foo(sli []int) {
    sli[0] = sli[0] + 1
}
func main() {
    a := []int{1, 2, 3}
    foo(a)
    fmt.Print(a)
}
```

### Variables as Functions

```go
var funcVar func(int) int

func incFn(x int) int {
    return x + 1
}

func main() {
    funcVar = incFn
    fmt.Print(funcVar(1))
}
```

###  Functions as Arguments

```go
func applyIt(afunct func (int) int, val int) int {
  return afunct(val)
}
```

- Example of using functions as arguments:

```go
func applyIt(afunct func (int) int, val int) int {
    return afunct(val)
}

func incFn(x int) int {return x + 1}

func decFn(x int) int {return x - 1}

func main() {
    fmt.Println(applyIt(incFn, 2))
    fmt.Println(applyIt(decFn, 2))
}
```

###  Anonymous Functions

```go
func applyIt(afunct func (int) int, val int) int {
    return afunct(val)
}

func main() {
    v := applyIt(
        func (x int) int { return x + 1 },
        2
    )
    fmt.Println(v)
}
```


###  Functions as Return Values

- Closure = function + its environment
- When functions are passed/returned, their environment comes with them!


```go
func MakeDistOrigin(o_x, o_y float64) func (float64, float64) float64 {
    fn := func (x, y float64) float64 {
        return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
    }
    return fn
}

func main() {
    Dist1 := MakeDistOrigin(0,0)
    Dist2 := MakeDistOrigin(2,2)
    fmt.Println(Dist1(2,2))
    fmt.Println(Dist2(2,2))
}
```

### Functions with Variable Argument Number

- Use ellipsis `...` to specify
- Treated as a slice inside function
- We can pass a slice to these functions as several arguments using `...` suffix

```go
func getMax(vals ...int) int {
    maxV := -1
    for _, v := range vals {
        if v > maxV {
            maxV = v
        }
    }
    return maxV
}

func main() {
    // Pass several arguments
    fmt.Println(getMax(1, 3, 6, 4))

    // Pass the same several arguments from slice
    vslice := []int{1, 3, 6, 4}
    fmt.Println(getMax(vslice...))
}
```


### Deferred Call Arguments

- Arguments of the Deferred Functionas are evaluated immediately, not in the deferred way.

```go
func main() {
    i := 1
    defer fmt.Println(i+1)  // i+1 is evaluated when program hits this line
                            // So, i = 1 and i + 1 = 2
                            // It is evaluated, however, not executed
                            // It will be 2, when it comes to the excution
    i++                     // Now i is incremented and i = 2
    fmt.Println(“Hello!”)
}

/*
Output will be:
Hello!
2
*/
```


## OOP

- Go can only hide data/methods in a package.
- Variables/functions are only exported if their names start with a capital letter.

```go
// (1) Defining type with methods:
type MyInt int

func (mi MyInt) Double () int {
    return int(mi*2)
}

// (2) Defining type struct with methods:
type Point struct {
    x float64
    y float64
}

func (p Point) DistToOrig() {
    t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
    return math.Sqrt(t)
}

func main() {
    // Using (1)
    v := MyInt(3)
    fmt.Println(v.Double())

    // Using (2)
    p1 := Point(3, 4)
    fmt.Println(p1.DistToOrig())
}
```

### Pointer Receivers

- Receiver is passed implicitly as an argument to the method.
- Receiver can be a pointer to a type.
- Call by reference, pointer is passed to the method.

 No Need to Reference:
- Point is referenced as `p`, not `*p`.
- Dereferencing is automatic with `.` operator.


```go
func (p *Point) OffsetX(v float64) {
    p.x = p.x + v
}
```


## Interfaces

- No need to explicitly state that function satisfies the interface.

```go
type Speaker interface {
    Speak ()
}

type Dog struct {
    name string
}

func (d Dog) Speak() {
    fmt.Println(d.name)
}

func main() {
    var s1 Speaker
    var d1 Dog{"Brian"}
    s1 = d1
    s1.Speak()
}
```


- **Empty interface** specifies no methods:

```go
func PrintMe(val interface{}) {
    fmt.Println(val)
}
```

### Dynamic Type vs Dynamic Value

- Dynamic Type: Concrete type which it is assigned to.
- Dynamic Value: Value of the dynamic type.

```go
type Speaker interface {
    Speak()
}

type Dog struct {
    name string
}

// func for Case 1
func (d Dog) Speak() {
    fmt.Println(d.name)
}

// func for Case 2
func (d *Dog) Speak() {
    if d == nil {
        fmt.Println("<noise>")
    } else {
        fmt.Println(d.name)
    }
}

func main() {
    // Case 1:
    // s1: Dynamic type is Dog, Dynamic value is d1
    var s1 Speaker
    var d1 Dog{“Brian”}
    s1 = d1
    s1.Speak()

    // Case 2:
    // d1 has no concrete value yet
    // s1 has a dynamic type but no dynamic value
    var s1 Speaker
    var d1 *Dog
    s1 = d1
    s1.Speak()
}
```

### Type Assertions

```go
// Type Assertions for Disambiguation
func DrawShape(s Shape2D) bool {
    rect, ok := s.(Rectangle)
    if ok {
        DrawRect(rect)
    }

    tri, ok := s.(Triangle)
    if ok {
        DrawTriangle(tri)
    }
}

// Type Switch
func DrawShape(s Shape2D) bool {
    switch sh := s.(type) {
    case Rectangle:
        DrawRect(sh)
    case Triangle:
        DrawTriangle(sh)
    }
}
```


## Error Handling

- Many Go programs return error interface objects to indicate errors.
- Correct operation: `error == nil`.
- Incorrect operation: `Error()` prints error message.

```go
type error interface {
    Error() string
}
```


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




