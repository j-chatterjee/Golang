//Write a Golang program to sum all the numbers provided in a given range.
package main

import (
	"fmt"
)

func main() {
	var start, end, sum int

	fmt.Print("Enter the start of the range: ")
	fmt.Scanln(&start)
	fmt.Print("Enter the end of the range: ")
	fmt.Scanln(&end)

	for i := start; i <= end; i++ {
		sum += i
	}

	fmt.Println("The sum of numbers in the range is:", sum)
}
