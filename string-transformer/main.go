// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Augustine Oche]
// Squad:  [The-Interface]

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var smallWords = map[string]bool{
	"a": true, "an": true, "the": true, "and": true, "but": true,
	"or": true, "for": true, "nor": true, "on": true, "at": true,
	"to": true, "by": true, "in": true, "of": true, "up": true,
	"as": true, "is": true, "it": true,
}

func toUpper(text string) string {
	return strings.ToUpper(text)
}

func toLower(text string) string {
	return strings.ToLower(text)
}

func toCap(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		w = strings.ToLower(w)
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

func toTitle(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		lower := strings.ToLower(w)
		if i == 0 || !smallWords[lower] {
			words[i] = strings.ToUpper(string(lower[0])) + lower[1:]
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}

func toSnake(text string) string {
	text = strings.ToLower(text)
	text = strings.Join(strings.Fields(text), "_")
	re := regexp.MustCompile(`[^a-z0-9_]`)
	return re.ReplaceAllString(text, "")
}

func toReverse(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		runes := []rune(w)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
	fmt.Println("──────────────────────────────────────")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Input: ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "exit" {
			fmt.Println("  Shutting down Augoboss String Transformer. Goodbye.")
			break
		}

		parts := strings.SplitN(line, " ", 2)
		cmd := strings.ToLower(parts[0])

		// Validate command first
		validCmds := "upper, lower, cap, title, snake, reverse, exit"
		validSet := map[string]bool{"upper": true, "lower": true, "cap": true, "title": true, "snake": true, "reverse": true}
		if !validSet[cmd] {
			fmt.Printf("  ✗ Unknown command: %q\n    Valid commands: %s\n\n", cmd, validCmds)
			continue
		}

		// Check for missing / whitespace-only text
		text := ""
		if len(parts) == 2 {
			text = strings.TrimSpace(parts[1])
			// Collapse internal spaces
			text = strings.Join(strings.Fields(text), " ")
		}
		if text == "" {
			fmt.Printf("  ✗ No text provided. Usage: %s <text>\n\n", cmd)
			continue
		}
		var result string
		switch cmd {
		case "upper":
			result = toUpper(text)
		case "lower":
			result = toLower(text)
		case "cap":
			result = toCap(text)
		case "title":
			result = toTitle(text)
		case "snake":
			result = toSnake(text)
		case "reverse":
			result = toReverse(text)
		}

		fmt.Printf("  → %s\n\n", result)
	}
}
