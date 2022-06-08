package main

import "fmt"

func main() {

	fmt.Println("Enter 1st number")
	var num1 int
	fmt.Scanln(&num1)
	var num2 int

	fmt.Println("Enter 2nd Number")
	fmt.Scanln(&num2)

	sum := num1 + num2
	fmt.Println("sum is :", sum)
}
