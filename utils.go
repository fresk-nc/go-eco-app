package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func clear() {
	print("\033[H\033[2J")
}

func printSeparator() {
	fmt.Println()
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println()
}

type strValidator func(s string) bool
type intValidator func(i int) bool

func readStringFromStdin(prompt string, isValid strValidator) string {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(prompt)
		text, err := reader.ReadString('\n')

		if err != nil || !isValid(strings.TrimSpace(text)) {
			fmt.Println("Ошибка ввода")
		} else {
			return strings.TrimSpace(text)
		}
	}
}

func readIntFromStdin(prompt string, isValid intValidator) int {
	for {
		var i int

		fmt.Print(prompt)

		_, err := fmt.Scanf("%d", &i)

		if err != nil || !isValid(i) {
			fmt.Println("Ошибка ввода")
		} else {
			return i
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
