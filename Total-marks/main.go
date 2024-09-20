//Write a Golang program to calculate a user's total marks for five subjects using user input
package main

import "fmt"

func main() {
	var marks [5]float64
	var total float64
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter marks for subject %d: ", i+1)
		fmt.Scanln(&marks[i])
		total += marks[i]
	}
	fmt.Println("Total marks obtained:", total)
}
