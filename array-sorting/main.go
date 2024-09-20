//Write a program to store random numbers in an array and find lower and highest values. Also, store value in an array where value will be sorted.
package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int

	fmt.Print("Enter the size of the array: ")
	fmt.Scanln(&size)
	numbers := make([]int, size)

	fmt.Println("Enter the elements:")
	for i := 0; i < size; i++ {
		fmt.Scanln(&numbers[i])
	}

	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	sort.Ints(numbers)

	fmt.Println("Minimum =", min)
	fmt.Println("Maximum =", max)
	fmt.Println("Sorted array =", numbers)
}
