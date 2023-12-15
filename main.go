package main

// create a new package
// in the main fcn, create a new slice of ints from 0 through 10
// iterate through the slice. for each ele, check to see if the num is even or odd
// if the value is even, print out "even" if it is odd, print out "odd"

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
