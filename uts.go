package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		for _, i := range text {
			fmt.Println(i)
		}
		fmt.Println()

		m := make(map[string]int)
		for _, i := range text {
			m[i]++
		}

		for key, element := range m {
			fmt.Printf("Kata %s : %d ", key, element)
			fmt.Println()
		}
	}
}
