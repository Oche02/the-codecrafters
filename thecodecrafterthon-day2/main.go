package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

func binToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}

func decimalTobinAndHex(dec string) string {
	decInt, _ := strconv.ParseInt(dec, 10, 64)
	toBin := strconv.FormatInt(decInt, 2)
	toHex := strconv.FormatInt(decInt, 16)
	fmt.Println("Binary: ", toBin)
	convert := strings.ToUpper(toHex)
	fmt.Println("Hexidecimal: ", convert)

	return ""
}

func main() {
	for {
		start := ""
		if start == "" {
			fmt.Println("=== Welcome to Augoboss CLI tool converter. press any key to continue ===")
		} else {
			fmt.Println("out of range")
			return
		}

		fmt.Println(start)
		fmt.Scanln(&start)

		fmt.Println("=== Welcome to Augoboss CLI tool converter ===")
		fmt.Println("1. Convert hexidecimal to desimal")
		fmt.Println("2. Convert binary to decimal")
		fmt.Println("3. convert decimal to binary and hexidecimal")
		fmt.Println("4. Help")
		fmt.Println("5. Exit")
		fmt.Println("choose from number 1 to 5")
		fmt.Println("")

		var choose string
		fmt.Scanln(&choose)
		switch choose {
		case "1":
			input := ("%d")
			fmt.Println("input Hexidecimal to convert")
			fmt.Scanln(&input)
			fmt.Print("Decimal: ")
			fmt.Println(hexToDecimal(input))
			fmt.Println("\n")
		case "2":
			input := ("%d")
			fmt.Println("input binary to convert")
			fmt.Scanln(&input)
			fmt.Print("Decimal: ")
			fmt.Println(binToDecimal(input))
			fmt.Println("\n")
		case "3":
			input := ("%d")
			fmt.Println("input decimal to convert")
			fmt.Scanln(&input)
			fmt.Println(decimalTobinAndHex(input))
			fmt.Println("\n")
		case "4":
			fmt.Println("==== Welcome to my Augoboss help center ====")
			fmt.Println("1. Press any key to continue after choosing my program to run")
			fmt.Println("2. choose an operator for your desired result")
			fmt.Println("3. type your desired input")
			fmt.Println("Enjoy Augoboss CLI tool converter....")
			fmt.Println("\n")
		case "5":
			fmt.Println("\n")
			fmt.Println("==== THANK YOU for using Augoboss CLI tool ====")
			fmt.Println("\n")
			return
		default:
			fmt.Println("\n")
			fmt.Println("Out of Range! input from 1 to 5")
			fmt.Println("\n")
			continue
		}
	}
}
