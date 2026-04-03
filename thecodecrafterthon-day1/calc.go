package main

import "fmt"

// "fmt"
// "strings"

func add(a, b int) int {
	return a+b
}

func subtract(a, b int) int {
	return a - b
}

func divide(a, b int) int {
	return a/b
}

func multiply(a, b int) int {
	return a*b
}

func main() {
	fmt.Println("Welcome CodeCrafter, this is a CLI calculator")
	for {
	fmt.Println("Choose an operation from the options below:")

	var opt int
	fmt.Println(" [1] addition\n [2] subtraction\n [3] Division\n [4] multiplication\n [5] Help\n [6] Quit")
	fmt.Scan(&opt)

	if opt <0 || opt > 6 {
		fmt.Println("invalid operation. Select [5] help for guidance")
		continue
		
	}

}
}