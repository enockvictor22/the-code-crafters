package main

import (
	"fmt"
)

func addition(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func divide(a, b float64) float64 {
	return a / b
}

func multiply(a, b float64) float64 {
	return a * b
}

func main() {
	fmt.Println("Welcome CodeCrafter, this is a CLI calculator")
	for {
		fmt.Println("Choose an operation from the options below:")

		var opt int
		fmt.Println(" [1] addition\n [2] subtraction\n [3] Division\n [4] multiplication\n [5] Help\n [6] Quit")
		fmt.Scan(&opt)

		if opt < 0 || opt > 6 {
			fmt.Println("invalid operation. Select [5] help for guidance")
			continue

		}
		var num1, num2 float64

		if opt == 6 {
			fmt.Println("Terminating Calculator, Goodbye...............")
			break
		}

		switch opt {
		case 1:
			fmt.Println("input first number")
			fmt.Scan(&num1)
			fmt.Println("input second number")
			fmt.Scan(&num2)
			fmt.Println("result: ", addition(num1, num2))
		case 2:
			fmt.Println("input first number")
			fmt.Scan(&num1)
			fmt.Println("input second number")
			fmt.Scan(&num2)
			fmt.Println("result: ", subtract(num1, num2))
		case 3:
			fmt.Println("input first number")
			fmt.Scan(&num1)
			fmt.Println("input second number")
			fmt.Scan(&num2)
			fmt.Println("result: ", divide(num1, num2))
		case 4:
			fmt.Println("input first number")
			fmt.Scan(&num1)
			fmt.Println("input second number")
			fmt.Scan(&num2)
			fmt.Println("result: ", multiply(num1, num2))
		case 5:
			fmt.Println("Guidance:")
			fmt.Println("select the number representing your \ndesired operation and input your \nnumber values for calculation")
		default:
			fmt.Println("please input a valid command")
		}

	}
}
