//Write a Golang program to print odd numbers in a given range.
package main

import (
	"fmt"
)

func main() {
	var start, end int

	fmt.Println("Enter the start of the range: ")
	fmt.Scanln(&start)
	fmt.Println("Enter the end of the range: ")
	fmt.Scanln(&end)

	fmt.Println("Odd numbers in the range:")

	for i := start; i <= end; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
