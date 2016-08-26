package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// A very special map
type SpecialMap struct {
	theMap map[int]string
	mutex  sync.RWMutex
}

func NewMap() SpecialMap {
	return SpecialMap{
		theMap: make(map[int]string),
	}
}

// puts a string to a random position in the map
func (sm *SpecialMap) Put(val string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sm.theMap[rand.Intn(80)] = val
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

}

// gets the string at position pos
func (sm *SpecialMap) Get(pos int) string {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	return sm.theMap[pos]
}

func (sm *SpecialMap) Print() {
	for key, val := range sm.theMap {
		fmt.Printf("%v: %v\n", key, val)
	}
}

func dontpanic() {
	if x := recover(); x != nil {
		fmt.Printf("I panicked.")
	}
}

func main() {
	defer dontpanic()

	rand.Seed(time.Now().UTC().UnixNano())

	m := NewMap()
	go func() {
		defer dontpanic()

		for i := 0; i < 1000; i++ {
			m.Put("hello")
		}
	}()

	go func() {
		defer dontpanic()

		for i := 0; i < 1000; i++ {
			m.Put("hi!")

		}
	}()

	time.Sleep(1000 * time.Millisecond)
	m.Print()
}
