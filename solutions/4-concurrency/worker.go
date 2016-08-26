package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func collectNoSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	//Timeout in go funcs?
	return c
}

func collect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case msg := <-input1:
				c <- msg
			case msg := <-input2:
				c <- msg
			case <-time.After(500 * time.Millisecond):
				c <- fmt.Sprintf("Timeout.")
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // comment out to have repoducilbe results
	c := collect(worker(1), worker(2))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done.")
}
