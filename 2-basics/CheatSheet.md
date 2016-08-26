# Cheat Sheet

## Variables

```go
// declaration & assignment
var x int
var s string = "foo"

// type inference
b := "bar"

// multiple assignment
var y, z int = 1, 2

// constants
const s string = "constant"
```


## Types
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

```

## Formatting

```go
print("hello")
fmt.Print("hello")
fmt.Println("hello")

fmt.Printf("Hello, %s", "Karlsruhe") fmt.Printf("Hello, %d", 42)
fmt.Printf("Hello, %v %v", "Karlsruhe", 47.11)
fmt.Printf("Hello, %+v", someStruct) // struct with fields
fmt.Printf("Hello, %T", someStruct) // type

```
See https://golang.org/pkg/fmt/

## If-Else

```go
x := 42
if x < 0 {
	fmt.Println("Negative")
} else if x == 42 {
	fmt.Println("The answer to life the universe and everything. Also positive.")
} else {
	fmt.Println("Positive")
}

// inline statement
if y := x % 2; y == 0 {
	fmt.Println("Even.")
} else {
	fmt.Println("Odd")
}
```

## Operators

```go
//comparison
== != < <= > >=

//logical
&&  ||  !

//arithmetical
*  /  % +  -

//bitwise
&  |  ^  &^  <<  >>

```

## Switch

```go
x := 4711
switch x {
case 4711:
	fmt.Println("Cologne")
case 42:
	fmt.Println("The answer to life.")
}


t := time.Now()
switch {
case t.Hour() >= 9 && t.Hour() < 12:
	fmt.Println("Good Morning")
case t.Hour() >= 12 && t.Hour() < 18:
	fmt.Println("Good Afternoon")
case t.Hour() >= 18 && t.Hour() < 22:
	fmt.Println("Good Evening")
default:
	fmt.Println("Good Night!")
}
```

## Arrays & Slices
```go
//arrays
var a [10]byte

b := [4]int {1, 2, 3, 4}
b[0] = 100

var x[5][100]

//slices
s := []int{3, 4, 2, 7, 5, 6, 10, 1, 9, 8}
s[0] = 100

x := s[1:] // all starting from second
y := s[0:len(s)-2] //all without last
z := s[1:3] // elements 2,3,4

t := make([]string, 0)
t = append(t, "foo")


```

## Loops

```go

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}


	for { // for ever
		fmt.Println(i)
	}


	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i, letter := range arr {
		fmt.Println(i, letter)
	}
```

## Functions

```go
func greet(name string) {
	fmt.Printf("Hello %s!\n", name)
}

func add(a int, b int) int {
	return a + b
}

func mult(a, b int) (c int) {
	c = a * b

	return c
}

//call functions
greet("World")
fmt.Println(mult(6, 7))
```

## Multi-Value Functions

```go
// two return values + combined parameters
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("%d is 0", b)
	}

	return a / b, nil
}

result, err := div(1, 0)
if err != nil {
	fmt.Println(err.Error())
}

if result, err := div(1, 0); err != nil {
	fmt.Println(err.Error())
}

```

## Functions as values
```go
func compute(a, b int, fn func(a, b int) int) int {
	return fn(a, b)
}

myFunc := func(a, b int) int {
	func1 := add
	func2 := mult
	return func2(a, func1(b, 1)) // a*(b+1)
}

fmt.Println(compute(7, 6, myFunc))

```

## Closures
```go
func makeGreeter(greeting string) func(string) string {
	count := 0

	return func(name string) string {
		count++
		return fmt.Sprintf("%s %s [%d times]", greeting, name, count)
	}
}

hello := makeGreeter("Hello")
whazzup := makeGreeter("Whazzup")

fmt.Println(hello("World"))        // => Hello World [1 times]
fmt.Println(whazzup("Karlsruhe"))  // => Whazzup Karlsruhe [1 times]
fmt.Println(whazzup("Hackschool")) // => Whazzup Hackschool [2 times]
```