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
