# Notes From Lectures

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

#### unicode package

```go
IsDigit(r rune)
IsSpace(r rune)
IsLetter(r rune)
IsLower(r rune)
IsPunct(r rune)

ToUpper(r rune)
ToLower(r rune)
```

#### strings package

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

#### strconv package

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
type struct Person {
    name string
    addr string
    phone string
}

var p1 Person

p1.name = "joe"
x = p1.addr

p1 := new(Person)  // sets all fields to "zero" values

p1 := Person(name: "joe", addr: "a st.", phone: "123")
```
