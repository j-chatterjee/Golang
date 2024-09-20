//Write a Golang program to check whether a user is eligible for voting using user input. If eligible output should 1, if not output should 0.
package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	eligible := age / 18
	fmt.Println(eligible)
}
