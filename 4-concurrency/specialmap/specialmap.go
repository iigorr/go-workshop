package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A very special map
type SpecialMap struct {
	theMap map[int]string
}

func NewMap() SpecialMap {
	return SpecialMap{
		theMap: make(map[int]string),
	}
}

// puts a string to a random position in the map
func (sm *SpecialMap) Put(val string) {
	sm.theMap[rand.Intn(80)] = val
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
}

// gets the string at position pos
func (sm *SpecialMap) Get(pos int) string {
	return sm.theMap[pos]
}

func (sm *SpecialMap) Print() {
	for key, val := range sm.theMap {
		fmt.Printf("%v: %v\n", key, val)
	}
}

func main() {
	m := NewMap()

	// TODO

	m.Print()
}
