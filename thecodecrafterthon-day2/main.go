package main

import (
	"fmt"
	"strconv"
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
	fmt.Println("Hexidecimal: ", toHex)
	return dec
}

func main() {
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
	fmt.Println("6. Help")
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
	case "2":
		input := ("%d")
		fmt.Println("input binary to convert")
		fmt.Scanln(&input)
		fmt.Print("Decimal: ")
		fmt.Println(binToDecimal(input))
	case "3":
		input := ("%d")
		fmt.Println("input decimal to convert")
		fmt.Scanln(&input)
		fmt.Println(decimalTobinAndHex(input))
	case "6":
		fmt.Println("1. choose an operator for your desired result")
	}
}
