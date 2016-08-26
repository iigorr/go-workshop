# Cheat Sheet

## goroutines

```go
go func() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, I am a go routine!")
}()
fmt.Println("Hello World!")

go gopher.Speak()

```

## Passing values to go routines
```go

values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
for _, n := range values {
	go func(val int) {
		 fmt.Println(val)
	 }(n)
}

```
## Channels
```go
	ch := make(chan Person)

	ch1 := make(chan int)    //unbuffered channel
	ch2 := make(chan int, 0) //unbuffered channel

	ch3 := make(chan *Person, 5) // buffered channel

	chan<- Person // send only
	<-chan Person // receive only
```

## Channels Sync
```go
ch := make(chan int)

go func() {
	doWork()
	ch <- 1
}()

doMoreWork()
<-ch
```

## Channels send data
```go
ch := make(chan string)

var a string
go func() {
	ch <- "hello"
}()

a = <-ch
fmt.Println(a)
```

## Example 1 - Worker
```go

doneCh := make(chan int)
personCh := make(chan Person)

go func(ch chan<- Person) {
	time.Sleep(1 * time.Second) //work
	ch <- Person{"Gopher", 4}   //send data
	time.Sleep(2 * time.Second) //work more
	doneCh <- 1                 // notify
}(personCh)

go func(ch <-chan Person) {
	p := <-ch
	fmt.Println(p)
}(personCh)

fmt.Println("Started workers. Waiting ...")
<-doneCh
fmt.Println("Done.")
```

## Example 2 - Semaphore
```go
var sem = make(chan int, MaxOutstanding)
func handle(r *Request) {
	sem <- 1 // Wait for active queue to drain.
	process(r) // May take a long time.
	<-sem // Done; enable next request to run.
}

func Serve(queue chan *Request) {
	for {
		req := <-queue
		go handle(req) // Don't wait for handle to finish.
	}
}

```

## Select

```go

ch1 := make(chan string)
ch2 := make(chan string)

select {
case s := <-ch1:
	fmt.Printf("Received string from channel 1: %v", s)
case s := <-ch2:
	fmt.Printf("Received string from channel 2: %v", s)
}

```

## Example - Ticker
```go
ticker := time.NewTicker(2 * time.Second)
personCh := make(chan Person)

go func() {
	for {
		select {
		case p := <-personCh:
			fmt.Printf("Received person %+v", p)
		case <-ticker.C:
			fmt.Printf("Received tick")
		}
	}
}()
time.Sleep(10 * time.Second)
personCh <- Person{"Gopher", 4}
```

## Mutex
```go
import "sync"

var mutex sync.Mutex
mutex.Lock()
mutex.Unlock()

var rwmutex sync.RWMutex
rwmutex.Lock()
rwmutex.Unlock()

rwmutex.RLock()
rwmutex.RUnlock()
```

## Mutex Example
```go
type TSInc struct {
	n   int
	mtx sync.RWMutex
}

func (ts *TSInc) Inc() {
	ts.mtx.Lock()
	ts.n = ts.n + 1
	ts.mtx.Unlock()
}

func (ts *TSInc) Val() int {
	ts.mtx.RLock()
	val := ts.n
	ts.mtx.RUnlock()

	return val
}

```

## defer Example
```go
func (ts *TSInc) Val() int {
	ts.mtx.RLock()
	defer ts.mtx.RUnlock()
		return ts.n
}

```

## panic & recover

```
defer func() {
	fmt.Printf("stopping...\n")
	if panicData := recover(); panicData != nil {
		fmt.Printf("Don't panic. Recovering from: %v\n", panicData)
	}
}()

panic("OMG OMG OMG!")

fmt.Printf("Hello!")
```