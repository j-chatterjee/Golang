//Write a program to store multiple user data from different server environments using the structure.
package main

import "fmt"

type User struct {
	Username    string
	Email       string
	Environment string
}

func main() {
	var users []User
	var numUsers int

	fmt.Print("Enter the number of users: ")
	fmt.Scanln(&numUsers)

	for i := 0; i < numUsers; i++ {
		var user User
		fmt.Printf("Enter details for user %d:\n", i+1)

		fmt.Print("Username: ")
		fmt.Scanln(&user.Username)

		fmt.Print("Email: ")
		fmt.Scanln(&user.Email)

		fmt.Print("Environment: ")
		fmt.Scanln(&user.Environment)

		users = append(users, user)
	}

	fmt.Println("\nUser Data:")
	for _, user := range users {
		fmt.Printf("Username: %s, Email: %s, Environment: %s\n", user.Username, user.Email, user.Environment)
	}
}
