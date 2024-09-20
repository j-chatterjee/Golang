//Write a program to create a map by taking two different arrays data.
package main

import "fmt"

func main() {

	var environments = make([]string, 4)
	var ec2Instances = make([]string, 4)

	fmt.Println("Enter environments (e.g., dev, qc, uat, prod):")
	for i := 0; i < 4; i++ {
		fmt.Printf("Environment %d: ", i+1)
		fmt.Scanln(&environments[i])
	}

	fmt.Println("Enter EC2 instances (e.g., t3.micro, t3.small, t3.medium, t3.large):")
	for i := 0; i < 4; i++ {
		fmt.Printf("EC2 Instance %d: ", i+1)
		fmt.Scanln(&ec2Instances[i])
	}

	instanceMap := make(map[string]string)
	for i := 0; i < 4; i++ {
		instanceMap[environments[i]] = ec2Instances[i]
	}

	fmt.Println("\nResulting Map:")
	for env, instance := range instanceMap {
		fmt.Printf("%s -> %s\n", env, instance)
	}
}
