package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %v", name)
}
