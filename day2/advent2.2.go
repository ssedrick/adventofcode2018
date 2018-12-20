package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)

	for scanner1.Scan() {
		line1 := scanner1.Text()
		// fmt.Printf("Comparing to %s\n", line1)

		file2, err := os.Open(filename)
		check(err)
		scanner2 := bufio.NewScanner(file2)
		for scanner2.Scan() {
			line2 := scanner2.Text()
			// fmt.Printf("line1: %s;\nline2: %s;\n", line1, line2)
			differs := 0

			line2Runes := []rune(line2)
			line1Runes := []rune(line1)
			for index, character1 := range line1Runes {
				character2 := line2Runes[index]
				if character2 != character1 {
					// fmt.Printf("%c != %c; ", character1, character2)
					differs++
				}
			}
			// fmt.Printf("differs by %d characters\n\n", differs)

			if differs == 1 {
				fmt.Printf("found: %s : %s\n", line1, line2)
			}
		}
		file2.Close()
	}
}
