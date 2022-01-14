# Notes From Lectures

### Functions

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

#### Call by value / reference

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
