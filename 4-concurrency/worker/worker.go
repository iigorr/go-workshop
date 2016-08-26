package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	c := collect(worker(1), worker(2))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done.")
}

func worker(id int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("Worker %d: %d", id, rand.Intn(100))
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}
	}()
	return ch
}

// TODO
