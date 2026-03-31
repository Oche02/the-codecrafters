package main

import (
	"fmt"
	"strings"
)

func cap(s string) string {
	return strings.Title(s)
}

func up(s string) string {
	return strings.ToUpper(s)
}

func low(s string) string {
	return strings.ToLower(s)
}

func main() {
	start := ""
	if start == "" {
		fmt.Println("==== Welcome To Augoboss String Transformer ====")
	} else {
		fmt.Println("")
	}
	fmt.Scanln(&start)

	fmt.Println("You made the right choice General!")
	fmt.Println("1. string")

	var choose string
	fmt.Scanln(&choose)
	switch choose {
	case "1":
		input := ("%s")
		fmt.Scanln(&input)
		fmt.Println(cap(input))
	}
}
