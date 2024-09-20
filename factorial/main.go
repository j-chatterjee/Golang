//Write a program to find out the factorial of a number using function and user input
package main

import "fmt"

func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var num int
    fmt.Print("Enter a number to find its factorial: ")
    fmt.Scan(&num)

    if num < 0 {
        fmt.Println("Factorial of a negative number is undefined.")
    } else {
        result := factorial(num)
        fmt.Printf("Factorial of %d is: %d\n", num, result)
    }
}
