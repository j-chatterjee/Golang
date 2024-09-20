//Write a Golang program to check a number even or odd using user input (No need to use if else logic, only need to print results in 0 and 1)
package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	fmt.Println((num + 1) % 2)
}
