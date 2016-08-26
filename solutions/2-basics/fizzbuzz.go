package main

import (
	"flag"
	"fmt"
)

const fizz = "Fizz"
const buzz = "Buzz"

func main() {
	limit := flag.Int("limit", 15, "How far should it go?")

	flag.Parse()

	for i := 1; i <= *limit; i++ {
		fizzbuzz(i)
	}

}

func fizzbuzz(num int) {
	if num%3 == 0 && num%5 == 0 {
		fmt.Println(fizz + buzz)
	} else if num%3 == 0 {
		fmt.Println(fizz)
	} else if num%5 == 0 {
		fmt.Println(buzz)
	} else {
		fmt.Println(num)
	}
}
