//Write a Golang program to calculate a user's total marks for five subjects using user input
package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Println(num1, "+", num2, "=", num1+num2)
	case "-":
		fmt.Println(num1, "-", num2, "=", num1-num2)
	case "*":
		fmt.Println(num1, "*", num2, "=", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println(num1, "/", num2, "=", num1/num2)
		} else {
			fmt.Println("Error: Division by zero!")
		}
	default:
		fmt.Println("Invalid input")
	}
}
