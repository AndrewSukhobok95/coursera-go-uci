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

#### Variables as Functions

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

####  Functions as Arguments

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

####  Anonymous Functions

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


####  Functions as Return Values

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

#### Functions with Variable Argument Number

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


#### Deferred Call Arguments

- Arguments of the Deferred Functionas are evaluated immediately, not in the deferred way.

```go
func main() {
    i := 1
    defer fmt.Println(i+1)  // i+1 is evaluated when program hits this line
                            // So, i = 1 and i + 1 = 2
    i++                     // Now i is incremented and i = 2
    fmt.Println(“Hello!”)
}

/*
Output will be:
Hello!
2
*/
```
