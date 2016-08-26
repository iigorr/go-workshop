# 2 basics

## Exercise 2.1 - FizzBuzz
Write a program that prints the numbers from 1 to 100. But for multiples of 3 print "Fizz" instead of the number and for the multiples of 5 print "Buzz". For numbers which are multiples of both 3 and 5 print "FizzBuzz".



### Expected results

1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
Fizz Buzz
...


### Variant 1
Use a switch statement

### Variant 2
Use a command line argument to pass the year

Hint: https://golang.org/pkg/flag/


## Exercise 2.2 - Bubble Sort
Implement bubble sort in GO.


```
func main() {
	arr := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", arr)
	bubbleSort(arr)
	fmt.Println("Sorted array: ", arr)
}
```

### Expected results
Sorting array  [3 4 2 7 5 6 10 1 9 8]
Done:  [1 2 3 4 5 6 7 8 9 10]


see http://www.tutorialspoint.com/data_structures_algorithms/bubble_sort_algorithm.htm

### Variant
Use a command line parameter to read an array length and generate a randomized sequence automatically.

Hint: See "flag" and "math/rand".


## Exercise 2.3 - Go Find


Implement the fibonacci series as a generator function:

```
func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
> 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
```

# Variant 1
Try a recursive implementation


## Exercise 2.4
Write a function `func gofind() (int, int, error)` that reads the file `effective-go.txt` and counts the number of lines that contain the strings "go" and "idiom" (case insensitively). Return both values or an error if something goes wrong.


Hints:
* Checkout `ioutil.ReadFile` from the "io/ioutil" package
* Checkout the "strings" package