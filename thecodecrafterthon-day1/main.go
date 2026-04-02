package main

import (
	"fmt"
)

func main() {
	for {

		var a, b int

		fmt.Println("===== Augoboss CLI Calculator =====")
		fmt.Println("           Input Number: ")

		fmt.Print("First input: ")
		firstInput, err := fmt.Scanf("%v\n", &a)
		if err != nil {
			fmt.Println("Error detected, first input", firstInput)
		}
		fmt.Print("Second input: ")
		secondInput, err := fmt.Scanf("%v\n", &b)
		if err != nil {
			fmt.Println("Error detected, second input", secondInput)

		}

		fmt.Println("==== Welcome To AUGOBOSS Calculator ====")
		fmt.Print("\n")
		fmt.Println("what can i do for you today?")
		fmt.Println("1. (+) Addition")
		fmt.Println("2. (-) Subtraction")
		fmt.Println("3. (*) Multiplicatin")
		fmt.Println("4. (/) Division")
		fmt.Println("5. Help")
		fmt.Println("6. Exit")
		fmt.Println("choose between 1 and 6")
		fmt.Print("Input: ")

		var choose string
		fmt.Scanln(&choose)
		switch choose {
		case "1":
			fmt.Println("Result: ", a+b)
		case "2":
			fmt.Println("Result: ", a-b)
		case "3":
			fmt.Println("Result: ", a*b)
		case "4":
			if b == 0 {
				fmt.Println("Error: must be divisable by a number ranging from (1) upward")
			} else {
				fmt.Println("Result: ", a/b)
			}
		case "5":
			fmt.Println("=== Welcome to Augoboss calculator help center ===")
			fmt.Println("How to make use of Augoboss calculator efficiently!")
			fmt.Println("1. Choose your first number you want the operator to work on")
			fmt.Println("2. choose your second number to operate upon")
			fmt.Println("3. Always remember to make a choice of an oprator for your calculation after choosing your both number")
			fmt.Println("          What is an OPRATOR?")
			fmt.Println("There are set of rules set to operate in you workspace or calculation to solve a particuler problems without any logical stress")
			fmt.Println("NOTE: in division, you can not divide any number by 0")
			fmt.Println("4. select Exit which is indicated in number (6) to exit out of the program")
		case "6":
			fmt.Println("     ==== Hope To See You Again ====")
			fmt.Println("Thank you for using Augoboss CLI Calculator")
			return
		default:
			fmt.Println("Out of range. Try again")
			continue
		}
	}
}
